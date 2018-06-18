package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"io"
	"io/ioutil"
    "github.com/gorilla/mux"
    "strconv" 
    "html/template"

)
func WebAppHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/recipes/app/webApp.html", http.StatusFound)
} 

//-------------------------JSON Handlers------------------------------//
func GetAllRecipesHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Get recipes request\n")
    recipes:=stored.GetRecipes()
    jsonSend, err := json.Marshal(recipes)
    errCatch(err)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonSend)
}

func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Create recipe request\n")
    var recipe Recipe 
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    errCatch(err)
    err = r.Body.Close()
    errCatch(err)
    err = json.Unmarshal(body, &recipe)
    if err != nil {
        w.WriteHeader(422) // unprocessable entity
        jsonerror,err:=json.Marshal(err)
        errCatch(err)
        w.Write(jsonerror)
    }

    recipe=stored.AddRecipe(recipe)
    jsonSend, err := json.Marshal(recipe)
    w.WriteHeader(http.StatusCreated)
    w.Write(jsonSend)
}

func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Read recipe request\n")

    vars := mux.Vars(r)

    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)

    recipe:=stored.ReadRecipe(recipeId)
    jsonSend, err := json.Marshal(recipe)
    errCatch(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonSend)
}

func EditRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Edit recipe request\n")
    var recipe Recipe 
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    vars := mux.Vars(r)
    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    errCatch(err)
    err = r.Body.Close()
    errCatch(err)
    err = json.Unmarshal(body, &recipe)

    if err != nil {
        w.WriteHeader(422) // unprocessable entity
        jsonerror,err:=json.Marshal(err)
        errCatch(err)
        w.Write(jsonerror)
        return
    }  
    recipe=stored.EditRecipe(recipeId, recipe)
    mresponse,err:=json.Marshal(recipe)
    errCatch(err)
    w.Write(mresponse)
}

func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Delete recipe request\n")

    vars := mux.Vars(r)
    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)
    msg:=stored.DeleteRecipe(recipeId)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if msg =="not found" {
        w.WriteHeader(422) // unprocessable entity
    }else{
        w.WriteHeader(http.StatusOK)
    }
    mresponse,err:=json.Marshal(msg)
    errCatch(err)
    w.Write(mresponse)
    
}

//--------------------------HTML Handlers---------------------------//

type RecipeList struct {
    Description string
    RecipesL []Recipe
}

func WebListHandler(w http.ResponseWriter, r *http.Request) {
    t := template.New("webList.html") 
    t, err := t.ParseFiles("assets/webList.html")
    errCatch(err)
    recipes:=stored.GetRecipes()
    d:=RecipeList{Description: "Recipe list: ", RecipesL:[]Recipe(recipes)}
    t.Execute(w, d)
}

func WebCreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
    t := template.New("webCreate.html") 
    t, err := t.ParseFiles("assets/webCreate.html")
    errCatch(err)
    d:=Recipe{};
    t.Execute(w, d)
} 

func WebGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    t := template.New("webView.html") 
    t, err := t.ParseFiles("assets/webView.html")
    errCatch(err)

    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)
    d:=stored.ReadRecipe(recipeId)
    t.Execute(w, d)
}

func WebDeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    t := template.New("webDelete.html") 
    t, err := t.ParseFiles("assets/webDelete.html")
    errCatch(err)

    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)

    msg:=stored.DeleteRecipe(recipeId)
    if msg =="not found" {
        w.WriteHeader(422) // unprocessable entity
    }else{
        w.WriteHeader(http.StatusOK)
    }

    d:="Recipe "+msg
    t.Execute(w, d)

}

func WebEditRecipeHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    t := template.New("webCreate.html") 
    t, err := t.ParseFiles("assets/webCreate.html")
    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)
    d:=stored.ReadRecipe(recipeId)
    t.Execute(w, d)
} 

