package main

import (
	srctrl "bindings_golang"
	"fmt"

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
	// fileId := srctrl.RecordFile(sourceFilePath)
	// srctrl.RecordFileLanguage(fileId, "golang") // doesn't support golang yet

	// sgi.Analyze(packagePath)
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

		cg_cha := cha.CallGraph(prog)

		for k, v := range cg_cha.Nodes {
			if k != nil && k.Name() != "init" {
				// logger.Println("CALLER:", k.Name())
				registerFuncBySSA(k)
				if len(v.Out) > 0 {
					for _, e := range v.Out {
						// logger.Println("CALLEE:", e.Callee.Func.Name())
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
