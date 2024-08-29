package dbreader

import (
	"encoding/xml"
	"os"
)

type XMLReader struct{}

func (x XMLReader) Read(filename string) (*Recipes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var recipes Recipes
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&recipes); err != nil {
		return nil, err
	}
	return &recipes, nil
}
