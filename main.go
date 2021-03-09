package main

import (
	"fmt"

	srctrl "github.com/CoatiSoftware/SourcetrailDB-master/build/bindings_golang"
)

func main() {
	prefix := "/home/tomatofaq/SourcetrailGolangIndexer/"
	databaseFilePath := prefix + "output/file.srctrldb"
	sourceFilePath := prefix + "example/file.py"

	if srctrl.Open(databaseFilePath) != true {
		fmt.Errorf("ERROR: {}", srctrl.GetLastError())
	}
	fmt.Println("Clearing loaded database now...")
	srctrl.Clear()

	srctrl.BeginTransaction()
	fileId := srctrl.RecordFile(sourceFilePath)
	srctrl.RecordFileLanguage(fileId, "python")

	srctrl.CommitTransaction()

	srctrl.Close()
}
