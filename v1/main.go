package main

import (
    "log"
    "net/http"
    "fmt"
)

func main() {
	AddRecipe(Recipe{Name:"Cooked rice",Servings:10,Tmins:20,
		IngNames:[]string{"Rice", "Salt"},
		IngAmounts:[]string{"500g", "20g"},
		Directions:[]string{"Cook","Eat"},})
	AddRecipe(Recipe{Name:"Lemonade",Servings:3,Tmins:5,
	IngNames:[]string{"Water", "Sugar"},
	IngAmounts:[]string{"0.5L", "60"},
	Directions:[]string{"Mix","Drink"},})

	router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}

func errCatch(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		panic(err)
	}
}