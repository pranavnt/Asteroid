package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func createCollection(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	os.Create("db/collections/" + name + ".properties")
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	collectionName := mux.Vars(r)["name"]
	filePath := "db/collections/" + collectionName + ".properties"

	fmt.Fprintf(w, filePath)
}
