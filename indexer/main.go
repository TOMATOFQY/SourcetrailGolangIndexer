package main

import (
	"flag"
	"os"
)

var indexer Indexer

func main() {
	path, _ := os.Getwd()
	pkgPath := path + "/../example"
	debug := false
	flag.StringVar(&pkgPath, "pkgPath", pkgPath, "The absolute path for target package. Redirect to the example folder by default.\n")
	flag.BoolVar(&debug, "debug", false, "Print log or not.")
	flag.Parse()

	indexer.DatabasePath = pkgPath + "/cg.srctrldb"
	indexer.Open()
	defer indexer.Close()
	cg(pkgPath)
}
