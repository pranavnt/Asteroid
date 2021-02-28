package main

import (
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
	docName := mux.Vars(r)["doc"]
	filePath = "db/collections/" + collectionName + ".properties"

}
