package main
import "fmt"
type Recipe struct {
    Id        int       `json:"rId"`
    Name      string    `json:"rname"`
    Servings   int       `json:"servings"`
    Tmins       int     `json:"tmins"`
    IngNames []string  `json:"ingNames"`
    IngAmounts []string  `json:"ingAmounts"`
    Directions []string  `json:"directions"`

}

type Recipes []Recipe

var currentId int

var recipesdb Recipes


func AddRecipe(recipe Recipe) Recipe {
    currentId += 1
    recipe.Id = currentId
    recipesdb = append(recipesdb, recipe)
    return recipe
}

func ReadRecipe(id int) Recipe {
    for _, v := range recipesdb {
        if v.Id == id {
            return v
        }
    }
    return Recipe{}
}

func EditRecipe(id int, newrecipe Recipe) Recipe {
    found:=0
    var ind int;
    for i, v := range recipesdb {
        if v.Id == id {
            //recipe= &v
            ind=i;
            found=1
            fmt.Printf("    recipe found")
        }
    }
    if found==1{
        if(newrecipe.Name!=""){
            recipesdb[ind].Name=newrecipe.Name
            fmt.Printf("    Name updated")
        }
        if(newrecipe.Servings!=0){
            recipesdb[ind].Servings=newrecipe.Servings
            fmt.Printf("    servings updated")
        }
        if(newrecipe.Tmins!=0){
            recipesdb[ind].Tmins=newrecipe.Tmins
            fmt.Printf("    tmins updated")
        }
        if(len(newrecipe.IngNames)!=0){
            recipesdb[ind].IngNames=newrecipe.IngNames
            fmt.Printf("    ing names updated")
        }
        if(len(newrecipe.IngAmounts)!=0){
            recipesdb[ind].IngAmounts=newrecipe.IngAmounts
            fmt.Printf("    ing amounts updated")
        }
        if(len(newrecipe.Directions)!=0){
            recipesdb[ind].Directions=newrecipe.Directions
            fmt.Printf("    directions updated")
        }

        return recipesdb[ind]
    }    
    return Recipe{}
}

func DeleteRecipe(id int) string {
    for i, v := range recipesdb {
        if v.Id == id {
            recipesdb = append(recipesdb[:i], recipesdb[i+1:]...)
            return "deleted"            
        }
    }
    return "not found"
}

