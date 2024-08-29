package main

import (
	"flag"
	"go-bootcamp/pkg/fscompare"
)

func main() {
	oldFilePath := flag.String("old", "", "path to the old file")
	newFilePath := flag.String("new", "", "path to the new file")

	flag.Parse()

	if *oldFilePath != "" && *newFilePath != "" {
		fscompare.CompareFS(*oldFilePath, *newFilePath)
	} else {
		println("Usage ./compareFS --old <fileOld> --new <fileNew>")
	}

}
