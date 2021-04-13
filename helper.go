package main

import (
	srctrl "bindings_golang"
	"encoding/json"
	"fmt"
)

func test() {
	if srctrl.Open(TestDatabaseFilePath) != true {
		fmt.Printf("ERROR: %#v\n", srctrl.GetLastError())
		return
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()

	nh1 := NameHierarchy{".", []NameElement{
		NameElement{"", "functionName", ""},
		NameElement{"", "instructionName1", ""}}}

	symbol, _ := json.Marshal(nh1)
	srctrl.RecordSymbol(string(symbol))

	nh2 := NameHierarchy{".", []NameElement{
		NameElement{"", "functionName", ""},
		NameElement{"", "instructionName2", ""}}}

	symbol, _ = json.Marshal(nh2)
	srctrl.RecordSymbol(string(symbol))

	// {
	// 	cfg := &packages.Config{
	// 		Mode:  packages.LoadAllSyntax,
	// 		Dir:   packagePath,
	// 		Tests: false,
	// 	}

	// 	initial, err := packages.Load(cfg)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	prog, _ := ssautil.AllPackages(initial, 0)
	// 	prog.Build()

	// 	cg_cha := cha.CallGraph(prog)

	// 	for k, v := range cg_cha.Nodes {
	// 		if k != nil && k.Name() != "init" {
	// 			// logger.Println("CALLER:", k.Name())
	// 			registerFuncBySSA(k)
	// 			if len(v.Out) > 0 {
	// 				for _, e := range v.Out {
	// 					// logger.Println("CALLEE:", e.Callee.Func.Name())
	// 					registerFuncBySSA(e.Callee.Func)
	// 					registerCallByEdge(e)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	srctrl.CommitTransaction()
	srctrl.Close()

}
