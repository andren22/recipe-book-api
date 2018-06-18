package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
	
	for _, staticfroute:= range staticFileRoutes{	
	staticFileDirectory := http.Dir("."+"/assets/")
	staticFileHandler := http.StripPrefix(staticfroute.Pattern, http.FileServer(staticFileDirectory))
	router.PathPrefix(staticfroute.Pattern).Handler(staticFileHandler).Methods("GET")
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