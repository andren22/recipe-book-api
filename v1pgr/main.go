package main

import (
    "log"
    "net/http"
    "fmt"
)
const (
	DB_HOST= "localhost"
	DB_PORT= 5432
    DB_USER     = "postgres"
    DB_PASSWORD = "1"
    DB_NAME     = "recipes_db"
)

func main() {


    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Printf("connection to databased failed")
		fmt.Scanln()
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping to database failed")
		fmt.Scanln()
		panic(err)
	}		
    errCatch(err)
    defer db.Close()

    initDB(db);
	router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}

func errCatch(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		panic(err)
	}
}