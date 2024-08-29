package dbcompare

import (
	"fmt"
	. "go-bootcamp/pkg/dbreader"
)

func CompareRecipes(oldRecipes, newRecipes *Recipes) {
	oldMap := make(map[string]Recipe)
	newMap := make(map[string]Recipe)

	// Заполнение карт для быстрого доступа
	for _, recipe := range oldRecipes.Cakes {
		oldMap[recipe.Name] = recipe
	}
	for _, recipe := range newRecipes.Cakes {
		newMap[recipe.Name] = recipe
	}

	// Найти удаленные и измененные рецепты
	for name, oldRecipe := range oldMap {
		if newRecipe, ok := newMap[name]; !ok {
			fmt.Println("REMOVED recipe:", name)
		} else {
			// Сравнить время приготовления
			if oldRecipe.StoveTime != newRecipe.StoveTime {
				fmt.Printf("CHANGED cooking time for %s - %s instead of %s\n", name, newRecipe.StoveTime, oldRecipe.StoveTime)
			}
			// Сравнить ингредиенты
			compareIngredients(oldRecipe.Ingredients, newRecipe.Ingredients)
		}
	}

	// Найти добавленные рецепты
	for name := range newMap {
		if _, ok := oldMap[name]; !ok {
			fmt.Println("ADDED recipe:", name)
		}
	}
}

// compareIngredients сравнивает два списка ингредиентов и выводит отчет о различиях
func compareIngredients(oldIngredients, newIngredients []Ingredient) {
	oldIngMap := make(map[string]Ingredient)
	newIngMap := make(map[string]Ingredient)

	for _, ing := range oldIngredients {
		oldIngMap[ing.Name] = ing
	}
	for _, ing := range newIngredients {
		newIngMap[ing.Name] = ing
	}

	for name, oldIng := range oldIngMap {
		if newIng, ok := newIngMap[name]; !ok {
			fmt.Println("REMOVED ingredient:", name)
		} else {
			if oldIng.Count != newIng.Count {
				fmt.Printf("CHANGED quantity of %s - %s instead of %s\n", name, newIng.Count, oldIng.Count)
			}
			if oldIng.Unit != newIng.Unit {
				fmt.Printf("CHANGED unit for %s - %s instead of %s\n", name, newIng.Unit, oldIng.Unit)
			}
		}
	}

	for name := range newIngMap {
		if _, ok := oldIngMap[name]; !ok {
			fmt.Println("ADDED ingredient:", name)
		}
	}
}
