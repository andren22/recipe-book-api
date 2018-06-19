# Recipes Book RESTful API

This repository contains an RESTful API written in Go that serves JSON and HTML content for creating, viewing, editing and deleting cooking recipes. Recipes data is stored locally by the server or in a PostgreSQL server database. A Javascript web application is also provided, allowing to directly use the API functionallity.

## Recipes content
Recipes implemented have the following content:
- ID: Number generated automatically by the server, this is an recipe unique identifier, used for view, delete and edit requests.
- Name: title given by the users.
- Yields: number of servings.
- Preparation time: minutes required for preparation.
- Ingredients: Unlimited ingredients, specified by name and amount.
- Directions: Unlimited instruction steps for preparation.

## API requests
### Javascript web application
The application is provided in response to `GET /recipes/app`

### JSON data
`GET json/recipes` Returns a JSON chain containing all the recipes stored in the server.

`POST json/recipes` Creates a recipe in the server with the JSON formated information provided in the body request.

`GET json/recipes/{ID}`Returns a JSON chain containing the information of the recipe with the ID number specified.

`DELETE json/recipes/{ID}` Deletes the recipe from the server with the specified ID.

`PUT json/recipes/{ID}` Reads the JSON formated data provided in the body request to update the recipe in the server with the specified ID number.

The JSON chain has the following format:  
```
[  
  {  
    "rId":(Recipe 1 ID),  
    "rname":"(Recipe name)",  
    "servings":(Recipe yields),  
    "tmins":(recipe preparation minutes: int),  
    "ingNames":["(Ingredient 1)","(Ingredient 2)",...],  
    "ingAmounts":["(Amount 1)","(Amount 2)",...],  
    "directions":["(Step 1)","(Step 2)"]  
  },  
  {  
    "rID":(Recipe 2 ID),  
    ...  
  },  
  ...  
]
```
### HTML pages
The API can respond request with rendered HTML pages using the following requests:  


`GET /web/recipes` Returns a web page displaying all the stored recipes.    
`GET /web/recipes/{ID}` Returns a web page displaying details of recipe with the ID number specified. Also, contains buttons for deleting and editing the recipe using Javascript functionality.  
`DELETE /web/recipes/{ID}` Returns a web page containing the response to the deletion of recipe with the ID number specified.  
`POST /web/recipes` Returns a web page containing an interface for recipe creation using Javascript.  
`PUT /web/recipes/{ID}` Returns a web page with an interface for editing the recipe with the ID number specified, through Javascript funcionality.  

## Using a database

The `/v1pgr` subfolder contains the API implementing a PostgreSQL database for data storage. To use the database, the table needed can be created using the following SQL commands:  
```
CREATE TABLE recipes  
(  
  id serial PRIMARY KEY,  
  name text NOT NULL,  
  servings int NOT NULL,  
  tmins int NOT NULL,  
  ingnames text[],  
  ingamounts text[],  
  directions text[]  
 );  
 ```
 The API configuration for database can be made at the beggining of the main.go file using the constants:  
    `DB_HOST` Host address.  
    `DB_PORT` TCP access port.  
    `DB_USER` Username.  
    `DB_PASSWORD` User password.  
    `DB_NAME` Database name.  
    
    
## Build and run

Before build, the following commands are required to get the required resources.  
`go get github.com/gorilla/mux`  
`go get github.com/lib/pq` 


Then, just build the executable by using:  
`go build`


## Future work
- Implement user registration and authentication for editing and deleting only recipes of his own.  
- Use unobtrusive JavaScript and reduce Javascript dependency of html pages.  
- Implement recipe rate system and most popular section.  
- Add support for MySQL databases.  


    
    
 
