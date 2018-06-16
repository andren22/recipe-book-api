
  
  function load(){

  var ingNamesList=[];
  var ingAmountsList=[];
  var directionsList=[];
  var currentId=0;
  showIndexScreen();
}

  function showCreationScreen(){
    sectionDisplay("indexScreen","hide");
    sectionDisplay("readScreen","hide");
    sectionDisplay("creationScreen","show");
    cleanFields();
    document.getElementById("createTitle").innerHTML="Create recipe: "; 
    document.getElementById("createb").innerHTML="Create recipe";
    document.getElementById("createb").setAttribute( "onClick", "javascript: sendRecipe();" );


  }

  function showEditScreen(){
    sectionDisplay("indexScreen","hide");
    sectionDisplay("readScreen","hide");
    sectionDisplay("creationScreen","show");
    document.getElementById("createTitle").innerHTML="Edit recipe: ";
    document.getElementById("createb").innerHTML="Edit recipe";
    document.getElementById("createb").setAttribute( "onClick", "javascript: sendEditRecipe();" );

  }

  function showIndexScreen(){
    cleanFields();
    loadRecipes();    
    sectionDisplay("creationScreen","hide");
    sectionDisplay("readScreen","hide");
    sectionDisplay("indexScreen","show");

  }

  function showReadScreen(){
    sectionDisplay("indexScreen","hide");
    sectionDisplay("creationScreen","hide");
    sectionDisplay("readScreen","show");

  }
  
  function sectionDisplay(sectionid,action) {        
      var x = document.getElementById(sectionid);
      if(action=="show") x.style.display = "block";
      if(action=="hide") x.style.display = "none";
  }

  function loadRecipes(){///////////MAIN METHOD
    document.getElementById("recipeslist").innerHTML="";
    recipeList = document.getElementById("recipeslist");

    fetch("/json/recipes")
      .then(response => response.json())
      .then(responseList => {
        responseList.forEach(recipe => {
          row = document.createElement("ul")
          rname = document.createElement("li")
          rlink=document.createElement("a")
          //rlink.href="#"
          rlink.setAttribute( "onClick", "javascript: showRecipe("+recipe.rId+");" );
          rlink.innerHTML=recipe.rname
          rdetails=document.createElement("p")
          rdetails.innerHTML=": "+recipe.servings+" servings, "+recipe.tmins+" minutes."          

          rlink.appendChild(rdetails)
          rname.appendChild(rlink)
          row.appendChild(rname)
          recipeList.appendChild(row)
        })
      }) 
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
      text = "<ul>";
      for (i = 0; i < iLen; i++) {
          text += "<li>" + ingNamesList[i] +": "+ingAmountsList[i]+ "</li>";
      }
        text += "</ul>";
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
    text = "<ol>";
    for (i = 0; i < iLen; i++) {
        text += "<li>" + directionsList[i] +"</li>";
    }
    text += "</ol>";  
    document.getElementById("input_dir").value="";
    document.getElementById("newDirections").innerHTML = text;
  }   
  

    function buildNewRecipe(){//////////////////////MAIN METHOD
      var recipe_name= document.getElementById("input_name").value;
      var recipe_yields=parseInt(document.getElementById("input_yields").value);      
      var recipe_tmins=parseInt(document.getElementById("input_tmins").value);
      if(recipe_name==""){alert("Write a recipe name"); return;}
      if(isNaN(recipe_yields) || recipe_yields==0){alert("Write a number in Yields field"); return;}
      if(isNaN(recipe_tmins) || recipe_tmins==0){alert("Write a number in Time field"); return;}
      if(ingNamesList.length==0){alert("Add at least one ingredient");return;}
      if(directionsList.length==0){alert("Add at least one direction");return;}
      

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
      .then(function(res){ console.log(res) })
      .catch(function(res){ console.log(res) })
      showIndexScreen();
  }

  function sendEditRecipe(){
      var url = "/json/recipes/"+currentId;
      new_recipe=buildNewRecipe();
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
      .then(function(res){ console.log(res) })
      .catch(function(res){ console.log(res) })
      showIndexScreen();
  }

      function cleanFields(){
        document.getElementById("input_name").value="";
        document.getElementById("input_yields").value="";
        document.getElementById("input_tmins").value="";
        document.getElementById("input_ing").value="";
        document.getElementById("input_amount").value="";
        document.getElementById("input_dir").value="";
        document.getElementById("newIngredients").innerHTML="";
        document.getElementById("newDirections").innerHTML="";
        document.getElementById("show_rname").innerHTML="";
        document.getElementById("show_yields").innerHTML="";
        document.getElementById("show_tmins").innerHTML="";
        document.getElementById("show_ingredients").innerHTML="";
        document.getElementById("show_directions").innerHTML="";
        ingNamesList=[];
        ingAmountsList=[];
        directionsList=[];
      }

      function showRecipe(rId){
        showReadScreen();
        currentId=rId;
        site="/json/recipes/"+rId;
        fetch(site)
          .then(response => response.json())
          .then(recipe => {
              document.getElementById("show_rname").innerHTML=recipe.rname;
              document.getElementById("show_yields").innerHTML=recipe.servings;
              document.getElementById("show_tmins").innerHTML=recipe.tmins;

              iLen = recipe.ingNames.length;
              text = "<ul>";
              for (i = 0; i < iLen; i++) {
                  text += "<li>" + recipe.ingNames[i] +": "+recipe.ingAmounts[i]+ "</li>";
              }
              text += "</ul>";
              document.getElementById("show_ingredients").innerHTML = text;
              
              iLen = recipe.directions.length;
              text = "<ol>";
              for (i = 0; i < iLen; i++) {
                  text += "<li>" + recipe.directions[i] +"</li>";
              }
              text += "</ol>";
              document.getElementById("show_directions").innerHTML = text;
          })
      }

     function fillEditRecipe(){
        showEditScreen();
        site="/json/recipes/"+currentId;
        fetch(site)
          .then(response => response.json())
          .then(recipe => {
              document.getElementById("input_name").value=recipe.rname;
              document.getElementById("input_yields").value=recipe.servings;
              document.getElementById("input_tmins").value=recipe.tmins;

              ingNamesList=recipe.ingNames;
              ingAmountsList=recipe.ingAmounts;
              directionsList=recipe.directions;
              displayDirs();
              displayIngs();              
          })
      }



    function deleteRecipe(){
      recipename=document.getElementById("show_rname").innerHTML;
      site="/json/recipes/"+currentId;
      if(confirm("Delete "+recipename+" recipe?")){
        var json = JSON.stringify({});

        fetch(site,
        {
            method: "DELETE"
        })
        .then(function(res){ console.log(res) })
        .catch(function(res){ console.log(res) })
        showIndexScreen();
      }
    }
