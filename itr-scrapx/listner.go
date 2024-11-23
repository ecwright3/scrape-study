package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func startListner() {
	//Create New Router
	r := mux.NewRouter()

	//Define a route
	r.Path("/mone/").Queries("target", "{key}").HandlerFunc(HomeHandler).Name("HomeHandler")

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	metadata, content := getImages(r.FormValue("target"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Add("Content-Disposition", "attachment; filename=Images.zip")

	w.Write(content.Bytes())

	data, err := json.Marshal(metadata)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	//Delete Archive from local storage
	err2 := os.Remove("Images.zip")

	if err2 != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

}
