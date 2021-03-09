package main

import (
	"fmt"

	srctrl "github.com/CoatiSoftware/SourcetrailDB-master/build/bindings_golang"
	sgi "github.com/tomatofaq/SourcetrailGolangIndexer/indexer"
)

const PREFIX string = "/home/tomatofaq/go/src/github.com/tomatofaq/SourcetrailGolangIndexer/"

func main() {
	databaseFilePath := PREFIX + "output/example.srctrldb"
	sourceFilePath := PREFIX + "example/main.go"
	package_path := PREFIX + "example/"

	sgi.Analyze(package_path)

	if srctrl.Open(databaseFilePath) != true {
		fmt.Printf("ERROR: %#v\n", srctrl.GetLastError())
		return
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()
	fileId := srctrl.RecordFile(sourceFilePath)
	srctrl.RecordFileLanguage(fileId, "golang") // doesn't support golang yet

	srctrl.CommitTransaction()

	srctrl.Close()
	fmt.Println(fileId)
}
