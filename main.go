package main

const PREFIX string = "/home/tomatofaq/go/src/github.com/tomatofaq/SourcetrailGolangIndexer/"
const packagePath = PREFIX + "example/"
const CGDatabaseFilePath = PREFIX + "output/cg.srctrldb"

var indexer Indexer = Indexer{DatabasePath: CGDatabaseFilePath}

func main() {
	cg(packagePath)
}
