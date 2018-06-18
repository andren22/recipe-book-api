//---------------View screen functions-------------//

function editPage(){
	rId=document.getElementById("rId").innerHTML;
  	site="/web/recipes/"+rId;
    fetch(site,
   	 {
    method: "PUT"
	  })
    .then(function(res){ return res.text(); })
    .then(function(res){ 
      document.body.innerHTML = res;
      init_edit();
     })
}

function deleteRecipe(){
  rId=document.getElementById("rId").innerHTML;
  site="/web/recipes/"+rId;
  if(confirm("Delete this recipe?")){
    var json = JSON.stringify({});

    fetch(site,
    {
        method: "DELETE"
    })
    .then(function(res){ return res.text();})
    .then(function(res){ document.body.innerHTML = res;})
     //alert("Recipe "+recipename+" deleted.");
     //document.location.href = "/web/recipes";
  }
}


//--------------Create page functions---------------//
function init_edit(){
  ingNamesList=[];
  ingAmountsList=[];
  directionsList=[];
  document.getElementById("createTitle").innerHTML="Edit recipe: ";
  document.getElementById("createb").innerHTML="Edit recipe";
  document.getElementById("createb").setAttribute( "onClick", "javascript: sendEditRecipe();" );
  readIngs();
  readDirs();
}
function init_create(){
  ingNamesList=[];
  ingAmountsList=[];
  directionsList=[];
  document.getElementById("createTitle").innerHTML="Create recipe: "; 
  document.getElementById("createb").innerHTML="Create recipe";
  document.getElementById("createb").setAttribute( "onClick", "javascript: sendRecipe();" );
}

function readIngs(){
  arr = document.getElementById("newIngredients").getElementsByTagName("li");
  for (i=0;i<arr.length;i++) {
    fields = arr[i].innerHTML.split(':');
    ingNamesList.push(fields[0]);
    ingAmountsList.push(fields[1]);
    }
}

function readDirs(){
  arr = document.getElementById("newDirections").getElementsByTagName("li");
  for (i=0;i<arr.length;i++) {
    directionsList.push(arr[i].innerHTML);
  }
}

function newIng() {
  ing_name=document.getElementById("input_ing").value;
  ing_amount=document.getElementById("input_amount").value;
  if (ing_name=="" || ing_amount==""){ window.alert("Please fill both fields");return;}
  ingNamesList.push(ing_name);
  ingAmountsList.push(ing_amount); 
  displayIngs();
} 
    
function displayIngs(){    
  iLen = ingNamesList.length;
  text = "";
  for (i = 0; i < iLen; i++) {
      text += "<li>" + ingNamesList[i] +": "+ingAmountsList[i]+ "</li>";
  }
    document.getElementById("input_ing").value="";
    document.getElementById("input_amount").value="";
    document.getElementById("newIngredients").innerHTML = text;  
} 

function newDir() {
  new_dir=document.getElementById("input_dir").value;
  if (new_dir==""){
    window.alert("Please write a direction");
  }
  else{      
    directionsList.push(new_dir);
  }
  displayDirs();
}

function deleteDir(){
  var num=directionsList.length;
  if(num==1){directionsList.splice(0,1);displayDirs();return;}
  if(num==0){alert("No directions to delete");return;}
  var dnum = parseInt(prompt("Insert step to delete",""));    
  if(isNaN(dnum)|| dnum==0 || dnum>num) {alert("Wrong value, insert a number between 1 and "+num);return;}
  directionsList.splice(dnum-1,1);
  displayDirs();
}

function displayDirs(){
  iLen = directionsList.length;
  text = "";
  for (i = 0; i < iLen; i++) {
      text += "<li>" + directionsList[i] +"</li>";
  } 
  document.getElementById("input_dir").value="";
  document.getElementById("newDirections").innerHTML = text;
}   

function buildNewRecipe(){
  var recipe_name= document.getElementById("input_name").value;
  var recipe_yields=parseInt(document.getElementById("input_yields").value);      
  var recipe_tmins=parseInt(document.getElementById("input_tmins").value);
  if(recipe_name==""){alert("Write a recipe name"); return false;}
  if(isNaN(recipe_yields) || recipe_yields==0){alert("Write a non-zero number in Yields field"); return false;}
  if(isNaN(recipe_tmins) || recipe_tmins==0){alert("Write a non-zero number in Time field"); return false;}
  if(ingNamesList.length==0){alert("Add at least one ingredient");return false;}
  if(directionsList.length==0){alert("Add at least one direction");return false;}
  var new_recipe={rname: recipe_name,
                  servings:recipe_yields,
                  tmins:recipe_tmins,
                  ingNames:ingNamesList,
                  ingAmounts: ingAmountsList,
                  directions: directionsList,
                  };
  return new_recipe;
}

function sendRecipe(){
    var url = "/json/recipes";
    new_recipe=buildNewRecipe();
    if(!new_recipe){return;}

    var json = JSON.stringify(new_recipe);
    
    fetch(url,
    {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        method: "POST",
        body: json
    })
    .then(function(res){ alert("Recipe "+new_recipe.rname+" created");goBack(); })
    .catch(function(res){ console.log(res) })
    
  
}

function sendEditRecipe(){
    rId=document.getElementById("rId").innerHTML;
    url = "/json/recipes/"+rId;
    new_recipe=buildNewRecipe();
    if(!new_recipe){return;}

    var json = JSON.stringify(new_recipe);
    
    fetch(url,
    {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        method: "PUT",
        body: json
    })
    .then(function(res){ alert("Recipe "+new_recipe.rname+" edited");goBack(); })
    .catch(function(res){ console.log(res) })
    showIndexScreen();
}

function goBack(){
  document.location.href = "/web/recipes";
}