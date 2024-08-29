package main

import (
	"flag"
	"fmt"

	. "go-bootcamp/pkg/dbcompare"
	"go-bootcamp/pkg/dbreader"
)

func main() {

	oldFile := flag.String("old", "", "Path to the old file")
	newFile := flag.String("new", "", "Path to the new file")

	flag.Parse()

	if *oldFile == "" && *newFile == "" {
		fmt.Println("Usage: ./compareDB --old <oldFileName> --new <newFileName>")
		return
	}

	recipeOld, err := dbreader.ReadRecipe(oldFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	recipeNew, err := dbreader.ReadRecipe(newFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	CompareRecipes(recipeOld, recipeNew)
}
