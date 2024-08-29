package dbreader

import (
	"encoding/json"
	"os"
)

type JSONReader struct{}

func (j JSONReader) Read(filename string) (*Recipes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var recipes Recipes
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&recipes); err != nil {
		return nil, err
	}
	return &recipes, nil
}
