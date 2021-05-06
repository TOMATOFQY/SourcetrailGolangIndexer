package main

import (
	srctrl "bindings_golang"
	"encoding/json"
	"go/ast"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

func registerFuncCFGbyFile(file *ast.File) {

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

	symbolId := srctrl.RecordSymbol(nh.string())
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
func (nh *NameHierarchy) pop() {
	nh.NameElements = nh.NameElements[:len(nh.NameElements)-1]
}

func (nh *NameHierarchy) string() string {
	ret, _ := json.Marshal(nh)
	return string(ret)
}

type NameElement struct {
	Prefix  string `json:"prefix"`
	Name    string `json:"name"`
	Postfix string `json:"postfix"`
}
