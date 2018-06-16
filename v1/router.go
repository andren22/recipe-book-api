package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
	
	for _, webroute:= range webroutes{	
	staticFileDirectory := http.Dir("."+"/assets/")
	staticFileHandler := http.StripPrefix(webroute.Pattern, http.FileServer(staticFileDirectory))
	router.PathPrefix(webroute.Pattern).Handler(staticFileHandler).Methods("GET")
	}

    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
}