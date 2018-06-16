package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
type WebRoute struct {
    Method      string
    Pattern     string
}

type Routes []Route
type WebRoutes []WebRoute

var routes = Routes{

    Route{
        "GetAllRecipes",
        "GET",
        "/json/recipes",
        GetAllRecipesHandler,
    },
    Route{
        "CreateRecipe",
        "POST",
        "/json/recipes",
        CreateRecipeHandler,
    },
	Route{
		"GetRecipe",
		"GET",
		"/json/recipes/{recipeId}",
		GetRecipeHandler,
	},
	Route{
		"EditRecipe",
		"PUT",
		"/json/recipes/{recipeId}",
		EditRecipeHandler,
	},
	Route{
		"DeleteRecipe",
		"DELETE",
		"/json/recipes/{recipeId}",
		DeleteRecipeHandler,
	},

    Route{
        "WebRecipeIndex",
        "GET",
        "/recipes",
        WebRecipeIndexHandler,
    },
}

var webroutes = WebRoutes{
    WebRoute{
        "GET",
        "/recipes/",
    },
}

