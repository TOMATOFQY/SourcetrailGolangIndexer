package main

import (
	srctrl "bindings_golang"
	"encoding/json"
	"fmt"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/ssa"
)

type Indexer struct {
	DatabasePath string
}

func (i Indexer) Open() error {
	if !srctrl.Open(i.DatabasePath) {
		return fmt.Errorf("ERROR: %#v", srctrl.GetLastError())
	}
	srctrl.Clear()
	return nil
}

func (i Indexer) BeginTransaction() {
	srctrl.BeginTransaction()
}

func (i Indexer) CommitTransaction() {
	srctrl.CommitTransaction()
}

func (i Indexer) Close() {
	srctrl.Close()
}

func (i Indexer) registerCallByEdge(e *callgraph.Edge) int {
	i.BeginTransaction()
	defer i.CommitTransaction()
	callerId := i.registerFunc(e.Caller.Func)
	calleeId := i.registerFunc(e.Callee.Func)
	referenceId := srctrl.RecordReference(callerId, calleeId, srctrl.REFERENCE_CALL)

	p := e.Site.Pos()
	prog := e.Callee.Func.Prog
	position := prog.Fset.Position(p)

	fileId := srctrl.RecordFile(position.Filename)
	srctrl.RecordReferenceLocation(referenceId, fileId, position.Line, position.Column-len(e.Callee.Func.Name()), position.Line, position.Column-1)
	return referenceId
}

func (i Indexer) registerFunc(f *ssa.Function) int {
	i.BeginTransaction()
	defer i.CommitTransaction()

	prog := f.Prog
	position := prog.Fset.Position(f.Pos())

	fileId := srctrl.RecordFile(position.Filename)
	srctrl.RecordFileLanguage(fileId, "cpp")

	pkg_name := ""
	if f.Pkg != nil && f.Pkg.Pkg != nil {
		pkg_name = f.Pkg.Pkg.Name()
	}

	Results := f.Signature.Results().String()
	if Results == "()" {
		Results = ""
	}

	nh := NameHierarchy{".", []NameElement{}}
	nh.push_back(NameElement{pkg_name, f.Name(), f.Signature.Params().String() + Results})

	symbolId := srctrl.RecordSymbol(nh.string())
	srctrl.RecordSymbolLocation(symbolId, fileId, position.Line, position.Column, position.Line, position.Column+len(f.Name())-1)
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
