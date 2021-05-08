package main

import (
	"log"
	"os"
)

const PREFIX string = "/home/tomatofaq/go/src/github.com/tomatofaq/SourcetrailGolangIndexer/"
const CFGDatabaseFilePath = PREFIX + "output/cfg.srctrldb"

const sourceFileName = "example/basicElement.go"
const sourceFilePath = PREFIX + sourceFileName

var logger = log.New(os.Stdout, "GLOBAL:\t", 0)

func main() {
	createCfg(sourceFilePath)
}
