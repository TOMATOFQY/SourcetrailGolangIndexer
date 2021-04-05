package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	srctrl "github.com/CoatiSoftware/SourcetrailDB-master/build/bindings_golang"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

const PREFIX string = "/home/tomatofaq/go/src/github.com/tomatofaq/SourcetrailGolangIndexer/"

var logger = log.New(os.Stdout, "GLOBAL:\t", 0)

func main() {

	databaseFilePath := PREFIX + "output/example.srctrldb"
	//	sourceFilePath := PREFIX + "example/main.go"
	package_path := PREFIX + "example/"

	if srctrl.Open(databaseFilePath) != true {
		fmt.Printf("ERROR: %#v\n", srctrl.GetLastError())
		return
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()
	// fileId := srctrl.RecordFile(sourceFilePath)
	// srctrl.RecordFileLanguage(fileId, "golang") // doesn't support golang yet

	// sgi.Analyze(package_path)
	{
		cfg := &packages.Config{
			Mode:  packages.LoadAllSyntax,
			Dir:   package_path,
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

func registerCallByEdge(e *callgraph.Edge) int {
	callerId := registerFuncBySSA(e.Caller.Func)
	calleeId := registerFuncBySSA(e.Callee.Func)
	referenceId := srctrl.RecordReference(callerId, calleeId, srctrl.REFERENCE_CALL)

	p := e.Site.Pos()
	prog := e.Callee.Func.Prog
	position := prog.Fset.Position(p)

	fileId := srctrl.RecordFile(position.Filename)
	srctrl.RecordReferenceLocation(referenceId, fileId, position.Line, position.Column-len(e.Callee.Func.Name()), position.Line, position.Column-1)
	return referenceId
}

func registerFuncBySSA(k *ssa.Function) int {
	prog := k.Prog
	position := prog.Fset.Position(k.Pos())

	fileId := srctrl.RecordFile(position.Filename)
	srctrl.RecordFileLanguage(fileId, "cpp")

	pkg_name := ""
	if k.Pkg != nil && k.Pkg.Pkg != nil {
		pkg_name = k.Pkg.Pkg.Name()
	}

	Results := k.Signature.Results().String()
	if Results == "()" {
		Results = ""
	}

	nh := NameHierarchy{".", []NameElement{}}
	nh.push_back(NameElement{pkg_name, k.Name(), k.Signature.Params().String() + Results})

	symbol, _ := json.Marshal(nh)
	symbolId := srctrl.RecordSymbol(string(symbol))
	srctrl.RecordSymbolLocation(symbolId, fileId, position.Line, position.Column, position.Line, position.Column+len(k.Name())-1)
	srctrl.RecordSymbolDefinitionKind(symbolId, srctrl.DEFINITION_EXPLICIT)
	srctrl.RecordSymbolKind(symbolId, srctrl.SYMBOL_FUNCTION)
	return symbolId
}

type NameHierarchy struct {
	NameDelimiter string        `json:"name_delimiter"`
	NameElements  []NameElement `json:"name_elements"`
}

func (nh *NameHierarchy) push_back(e NameElement) {
	nh.NameElements = append(nh.NameElements, e)
}

type NameElement struct {
	Prefix  string `json:"prefix"`
	Name    string `json:"name"`
	Postfix string `json:"postfix"`
}
