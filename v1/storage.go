package main
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
            ind=i;
            found=1
        }
    }
    if found==1{
        if(newrecipe.Name!=""){
            recipesdb[ind].Name=newrecipe.Name
        }
        if(newrecipe.Servings!=0){
            recipesdb[ind].Servings=newrecipe.Servings
        }
        if(newrecipe.Tmins!=0){
            recipesdb[ind].Tmins=newrecipe.Tmins
        }
        if(len(newrecipe.IngNames)!=0){
            recipesdb[ind].IngNames=newrecipe.IngNames
        }
        if(len(newrecipe.IngAmounts)!=0){
            recipesdb[ind].IngAmounts=newrecipe.IngAmounts
        }
        if(len(newrecipe.Directions)!=0){
            recipesdb[ind].Directions=newrecipe.Directions
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

