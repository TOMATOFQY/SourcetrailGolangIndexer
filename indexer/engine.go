package sourcetrailgolangindexer

import (
	"fmt"

	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
)

func init() {
}

func Analyze(packagePath string) {
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
		if k != nil {
			fmt.Println("CALLER:", k.Name())
			fmt.Println(prog.Fset.Position(k.Pos()))
			if len(v.Out) > 0 {
				fmt.Println(v.Out)
				for _, e := range v.Out {
					p := e.Site.Pos()
					fmt.Println(prog.Fset.Position(p))
				}
			}
			fmt.Println()
		}
	}
}
