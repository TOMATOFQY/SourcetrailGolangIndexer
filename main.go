package main

import (
	"log"
	"os"
)

const PREFIX string = "/home/tomatofaq/go/src/github.com/tomatofaq/SourcetrailGolangIndexer/"
const packagePath = PREFIX + "example/"
const CGDatabaseFilePath = PREFIX + "output/cg.srctrldb"

var logger = log.New(os.Stdout, "GLOBAL:\t", 0)

func main() {
	cg(packagePath)
}
