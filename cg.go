package main

import (
	srctrl "bindings_golang"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func cg(packagePath string) {
	if srctrl.Open(CGDatabaseFilePath) != true {
		fmt.Printf("ERROR: %#v\n", srctrl.GetLastError())
		return
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()
	{
		cfg := &packages.Config{
			Mode:  packages.LoadAllSyntax,
			Dir:   packagePath,
			Tests: false,
		}

		initial, err := packages.Load(cfg)
		if err != nil {
			panic(err)
		}

		prog, _ := ssautil.AllPackages(initial, 0)
		prog.Build()

		cg := cha.CallGraph(prog)

		ast.Print(nil, cg)

		for k, v := range cg.Nodes {
			if k != nil && k.Name() != "init" {
				registerFuncBySSA(k)
				if len(v.Out) > 0 {
					for _, e := range v.Out {
						registerFuncBySSA(e.Callee.Func)
						registerCallByEdge(e)
					}
				}
			}
		}
	}
	srctrl.CommitTransaction()
	srctrl.Close()
}
