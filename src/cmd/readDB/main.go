package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"

	"go-bootcamp/pkg/dbreader"
	"os"
	"strings"
)

func main() {

	filename := flag.String("f", "", "Path to the file")
	flag.Parse()

	if len(os.Args) != 3 && *filename == "" {
		fmt.Println("Usage: ./readDB -f <filename>")
		return
	}

	var reader dbreader.RecipeReader

	if strings.HasSuffix(*filename, ".xml") {
		reader = dbreader.XMLReader{}
	} else if strings.HasSuffix(*filename, ".json") {
		reader = dbreader.JSONReader{}
	} else {
		fmt.Println("Unsupported file format")
		return
	}

	recipes, err := reader.Read(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if strings.HasSuffix(*filename, ".xml") {
		jsonData, err := json.MarshalIndent(recipes, "", "    ")
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		xmlData, err := xml.MarshalIndent(recipes, "", "    ")
		if err != nil {
			fmt.Println("Error converting to XML:", err)
			return
		}
		fmt.Println(string(xmlData))
	}
}
