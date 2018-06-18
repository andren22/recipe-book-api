package main

import (
    "database/sql"
    "github.com/lib/pq"
    "fmt"
    )
type Recipe struct {
    Id        int       `json:"rId"`
    Name      string    `json:"rname"`
    Servings   int       `json:"servings"`
    Tmins       int     `json:"tmins"`
    IngNames []string  `json:"ingNames"`
    IngAmounts []string  `json:"ingAmounts"`
    Directions []string  `json:"directions"`

}

type DBstored struct {
    db *sql.DB
}

var stored DBstored

func InitDB(s *sql.DB) {
    //store = s
    stored=DBstored{s}
    stored.AddRecipe(Recipe{Name:"Cooked rice",Servings:10,Tmins:20,
        IngNames:[]string{"Rice", "Salt"},
        IngAmounts:[]string{"500g", "20g"},
        Directions:[]string{"Cook","Eat"},})
    stored.AddRecipe(Recipe{Name:"Lemonade",Servings:3,Tmins:5,
        IngNames:[]string{"Water", "Sugar"},
        IngAmounts:[]string{"0.5L", "60"},
        Directions:[]string{"Mix","Drink"},})
}
func (store *DBstored) GetRecipes() ([]Recipe) {

    rows, err := store.db.Query("SELECT * from recipes")
    if err != nil { return []Recipe{}}
    defer rows.Close()

    recipes := []Recipe{}
    for rows.Next() {
        recipe := &Recipe{}
         if err := rows.Scan(
                &recipe.Id,
                &recipe.Name,
                &recipe.Servings,
                &recipe.Tmins,
                pq.Array(&recipe.IngNames),
                pq.Array(&recipe.IngAmounts),
                pq.Array(&recipe.Directions),
                ); 
            err != nil { errCatch(err) }        
        recipes = append(recipes, *recipe)
    }
    return recipes
}

func (received *DBstored) AddRecipe(recipe Recipe) Recipe {

    var lastInsertId int
    squery:="insert into recipes(name,servings,tmins,ingnames,ingamounts,directions) values($1,$2,$3,$4,$5,$6) returning id"
    err:= received.db.QueryRow(squery, 
        recipe.Name,
        recipe.Servings,
        recipe.Tmins,
        pq.Array(recipe.IngNames),
        pq.Array(recipe.IngAmounts),
        pq.Array(recipe.Directions),
        ).Scan(&lastInsertId)
    errCatch(err)
    recipe.Id=lastInsertId
    return recipe
}

func (received *DBstored) ReadRecipe(id int) Recipe {
    rows, err := received.db.Query("select * from recipes where id=$1",id)
    if err != nil { return Recipe{} }
    defer rows.Close()

    recipe := &Recipe{}
    rows.Next()
    if err := rows.Scan(
        &recipe.Id,
        &recipe.Name,
        &recipe.Servings,
        &recipe.Tmins,
        pq.Array(&recipe.IngNames),
        pq.Array(&recipe.IngAmounts),
        pq.Array(&recipe.Directions),
        ); 
    err != nil {errCatch(err)}
    
    return *recipe
}

func (received *DBstored) EditRecipe(id int, newrecipe Recipe) Recipe {
    var exist bool
    err:= received.db.QueryRow("select exists(select * from recipes where id=$1)",id).Scan(&exist)
    errCatch(err)
    if exist!=true{ 
        fmt.Printf("delete recipe not found")
        return Recipe{}
    }

    if(newrecipe.Name!=""){
        _,err:=received.db.Query("update recipes set name=$1 where id=$2",newrecipe.Name,id)
        errCatch(err)
    }

    if(newrecipe.Servings!=0){
        _,err:=received.db.Query("update recipes set servings=$1 where id=$2",newrecipe.Servings,id)
        errCatch(err)
    }

    if(newrecipe.Tmins!=0){
        _,err:=received.db.Query("update recipes set tmins=$1 where id=$2",newrecipe.Tmins,id)
        errCatch(err)
    }

    if(len(newrecipe.IngNames)!=0){
        _,err:=received.db.Query("update recipes set ingnames=$1 where id=$2",pq.Array(newrecipe.IngNames),id)
        errCatch(err)
    }

    if(len(newrecipe.IngAmounts)!=0){
        _,err:=received.db.Query("update recipes set ingamounts=$1 where id=$2",pq.Array(newrecipe.IngAmounts),id)
        errCatch(err)
    }

    if(len(newrecipe.Directions)!=0){
        _,err:=received.db.Query("update recipes set directions=$1 where id=$2",pq.Array(newrecipe.Directions),id)
        errCatch(err)
    }
    return Recipe{}
}

func (received *DBstored) DeleteRecipe(id int) string {
    var exist bool
    err:= received.db.QueryRow("select exists(select * from recipes where id=$1)",id).Scan(&exist)
    errCatch(err)
    if exist!=true{ 
        fmt.Printf("delete recipe not found")
        return "not found"
    }
    _, err=received.db.Query("delete from recipes where id=$1",id)
    errCatch(err)

    return "deleted"           
}

