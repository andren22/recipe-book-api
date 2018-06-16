package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {
	AddRecipe(Recipe{Name:"First",Servings:5,Tmins:30,
		IngNames:[]string{"rize", "salt"},
		IngAmounts:[]string{"500g", "30g"},
		Directions:[]string{"cook","eat"},})
	AddRecipe(Recipe{Name:"Second",Servings:5,Tmins:30,
	IngNames:[]string{"water", "sugar"},
	IngAmounts:[]string{"0.5L", "90g"},
	Directions:[]string{"cook","eat"},})

	router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}

func errCatch(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		panic(err)
	}
}