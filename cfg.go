package main

import (
	srctrl "bindings_golang"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"

	"golang.org/x/tools/go/cfg"
)

var (
	fset   *token.FileSet
	lines  []string
	fileId int
)

func createCfg(filePath string) {
	if srctrl.Open(CFGDatabaseFilePath) != true {
		fmt.Printf("ERROR: %#v\n", srctrl.GetLastError())
		return
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()
	{
		fileId = srctrl.RecordFile(filePath)
		srctrl.RecordFileLanguage(fileId, "cpp")
		src, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Errorf("ERROR！")
		}

		lines = strings.Split(string(src), "\n")
		fset = token.NewFileSet()
		f, _ := parser.ParseFile(fset, filePath, src, parser.Mode(0))

		for _, decl := range f.Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok {
				registerFuncCFG(decl)
			}
		}
	}
	srctrl.CommitTransaction()
	srctrl.Close()
}

func registerFuncCFG(funcDecl *ast.FuncDecl) {
	g := cfg.New(funcDecl.Body, mayReturn)
	for _, b := range g.Blocks {
		// 记录单个block之内的指令
		var prev *int
		for _, node := range b.Nodes {
			position := fset.Position(node.Pos())
			symbolId := registerInstruction(funcDecl, b, &node)
			srctrl.RecordSymbolLocation(symbolId, fileId, position.Line, position.Column, position.Line, position.Column)
			if prev != nil {
				referenceId := srctrl.RecordReference(*prev, symbolId, srctrl.REFERENCE_CALL)
				srctrl.RecordReferenceLocation(referenceId, fileId, position.Line, position.Column, position.Line, position.Column)
			}
			prev = &symbolId
		}

		// 记录block之间的流向
		for _, nb := range b.Succs {
			linkBlocks(funcDecl, b, nb)
		}

	}
}

func linkBlocks(funcDecl *ast.FuncDecl, b *cfg.Block, nb *cfg.Block) int {
	bId := registerBlock(funcDecl, b)
	nbId := registerBlock(funcDecl, nb)
	lb, lnb := len(b.Nodes), len(nb.Nodes)
	var referenceId int
	var pos token.Pos
	if lb > 0 && lnb > 0 {
		tail := registerInstruction(funcDecl, b, &b.Nodes[len(b.Nodes)-1])
		head := registerInstruction(funcDecl, nb, &nb.Nodes[0])
		referenceId = srctrl.RecordReference(tail, head, srctrl.REFERENCE_CALL)
		pos = b.Nodes[len(b.Nodes)-1].Pos()
	} else if lb > 0 {
		tail := registerInstruction(funcDecl, b, &b.Nodes[len(b.Nodes)-1])
		referenceId = srctrl.RecordReference(tail, nbId, srctrl.REFERENCE_CALL)
		pos = b.Nodes[len(b.Nodes)-1].Pos()
	} else if lnb > 0 {
		head := registerInstruction(funcDecl, nb, &nb.Nodes[0])
		referenceId = srctrl.RecordReference(bId, head, srctrl.REFERENCE_CALL)
		pos = nb.Nodes[0].Pos()
	} else {
		referenceId = srctrl.RecordReference(bId, nbId, srctrl.REFERENCE_CALL)
	}

	position := fset.Position(pos)
	srctrl.RecordReferenceLocation(referenceId, fileId, position.Line, position.Column, position.Line, position.Column)
	return referenceId
}

func registerBlock(funcDecl *ast.FuncDecl, b *cfg.Block) int {
	nh := NameHierarchy{".", []NameElement{}}
	nh.push_back(NameElement{"", funcDecl.Name.Name + " " + b.String(), ""})
	symbolId := srctrl.RecordSymbol(nh.string())
	return symbolId
}

func registerInstruction(funcDecl *ast.FuncDecl, b *cfg.Block, node *ast.Node) int {
	nh := NameHierarchy{".", []NameElement{}}
	nh.push_back(NameElement{"", funcDecl.Name.Name + " " + b.String(), ""})
	position := fset.Position((*node).Pos())
	line := lines[position.Line-1]
	line = strings.TrimLeft(line, "\t")
	nh.push_back(NameElement{"", fmt.Sprintf("%v", position.Line) + " " + line, ""})
	symbolId := srctrl.RecordSymbol(nh.string())
	srctrl.RecordSymbolKind(symbolId, srctrl.SYMBOL_FUNCTION)
	return symbolId
}

func mayReturn(call *ast.CallExpr) bool {
	switch fun := call.Fun.(type) {
	case *ast.Ident:
		return fun.Name != "panic"
	case *ast.SelectorExpr:
		return fun.Sel.Name != "Fatal"
	}
	return true
}
