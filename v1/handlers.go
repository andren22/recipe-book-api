package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"io"
	"io/ioutil"
    "github.com/gorilla/mux"
    "strconv" 

)
func WebRecipesAppHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/recipes/app/webApp.html", http.StatusFound)
} 

func GetAllRecipesHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Get recipes request\n")
    jsonSend, err := json.Marshal(recipesdb)
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

    recipe=AddRecipe(recipe)
    jsonSend, err := json.Marshal(recipe)
    w.WriteHeader(http.StatusCreated)
    w.Write(jsonSend)
}

func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Read recipe request\n")

    vars := mux.Vars(r)

    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)

    recipe:=ReadRecipe(recipeId)
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
    recipe=EditRecipe(recipeId, recipe)
    mresponse,err:=json.Marshal(recipe)
    errCatch(err)
    w.Write(mresponse)
}

func DeleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Delete recipe request\n")

    vars := mux.Vars(r)
    recipeId, err := strconv.Atoi(vars["recipeId"])
    errCatch(err)
    msg:=DeleteRecipe(recipeId)
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


