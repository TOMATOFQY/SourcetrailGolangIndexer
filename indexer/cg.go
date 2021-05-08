package main

import (
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func cg(packagePath string) {

	prog := createProg(packagePath)

	cg := cha.CallGraph(prog)
	for k, v := range cg.Nodes {
		if k != nil && k.Name() != "init" {
			indexer.registerFunc(k)
			if len(v.Out) > 0 {
				for _, e := range v.Out {
					indexer.registerFunc(e.Callee.Func)
					indexer.registerCallByEdge(e)
				}
			}
		}
	}
}

func createProg(packagePath string) *ssa.Program {
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

	return prog
}
