package dbreader

import (
	"errors"
	"fmt"
	"strings"
)

type RecipeReader interface {
	Read(file string) (*Recipes, error)
}

type Recipes struct {
	Cakes []Recipe `xml:"cake" json:"cake"`
}

type Recipe struct {
	Name        string       `xml:"name" json:"name"`
	StoveTime   string       `xml:"stovetime" json:"time"`
	Ingredients []Ingredient `xml:"ingredients>item" json:"ingredients"`
}

type Ingredient struct {
	Name  string `xml:"itemname" json:"ingredient_name"`
	Count string `xml:"itemcount" json:"ingredient_count"`
	Unit  string `xml:"itemunit,omitempty" json:"ingredient_unit,omitempty"`
}

func ReadRecipe(filename *string) (*Recipes, error) {

	var reader RecipeReader

	if strings.HasSuffix(*filename, ".xml") {
		reader = XMLReader{}
	} else if strings.HasSuffix(*filename, ".json") {
		reader = JSONReader{}
	} else {
		fmt.Println("Unsupported file format")
		return nil, errors.New("unsupported file format")

	}

	recipes, err := reader.Read(*filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return recipes, nil
}
