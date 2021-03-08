package main

import (
	"fmt"

	srctrldb "github.com/CoatiSoftware/SourcetrailDB-master/build/bindings_golang"
)

func main() {
	t := srctrldb.IsEmpty()
	fmt.Println(t)
}

