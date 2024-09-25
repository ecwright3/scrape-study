package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func startListner() {
	//Create New Router
	r := mux.NewRouter()

	//Define a route
	r.Path("/mone/").Queries("target", "{key}").HandlerFunc(HomeHandler).Name("HomeHandler")
	//r.HandleFunc("/{name}", HomeHandler)
	//http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	getImages(r.FormValue("target"))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Name: %v\n", r.FormValue("target"))

}
