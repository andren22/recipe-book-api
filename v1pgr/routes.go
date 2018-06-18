package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
type StaticRoute struct {
    Method      string
    Pattern     string
}

type Routes []Route
type StaticRoutes []StaticRoute

var routes = Routes{
    //--------JSON Routes-------//
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
    //-------Web App Route-------//
    Route{
        "WebRecipeApp",
        "GET",
        "/recipes/app",
        WebAppHandler,
    },

    //-------HTML Routes---------//
    Route{
        "WebRecipeList",
        "GET",
        "/web/recipes",
        WebListHandler,
    },
    Route{
        "WebRecipeList",
        "POST",
        "/web/recipes",
        WebCreateRecipeHandler,
    },
    
    Route{
        "WebShowRecipe",
        "GET",
        "/web/recipes/{recipeId}",
        WebGetRecipeHandler,
    },

    Route{
        "WebDeleteRecipe",
        "DELETE",
        "/web/recipes/{recipeId}",
        WebDeleteRecipeHandler,
    },

    Route{
        "WebEditRecipe",
        "PUT",
        "/web/recipes/{recipeId}",
        WebEditRecipeHandler,
    },


}

var staticFileRoutes = StaticRoutes{
    StaticRoute{
        "GET",
        "/recipes/app/",
    },

    StaticRoute{
        "GET",
        "/styles/",
    },

    StaticRoute{
        "GET",
        "/web/recipes/create/",
    },
}

