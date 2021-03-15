package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pranavnt/mamba"
)

func main() {
	app := mamba.New()
	app.AddCommand("start", start)
	app.AddCommand("delete", deleteCollection)
	app.Run(os.Args)
}

func start(params mamba.Dict) {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/login", login)
	r.HandleFunc("/signUp", signUp)

	// CREATE operations
	r.HandleFunc("/api/collection/{name}", createCollection)
	r.HandleFunc("/api/collection/{name}/create/document", createDocument)

	// READ operations
	r.HandleFunc("/api/collection/{name}/document/read/{doc}", readDocument)
	r.HandleFunc("/api/collection/{name}/document/read/user/{uid}", readDocumentsById)
	r.HandleFunc("/api/collection/{name}/all", readAllDocuments)

	// UPDATE operation
	r.HandleFunc("/api/collection/{name}/document/update/{doc}", updateDocument)

	// DELETE operation
	r.HandleFunc("/api/collection/{name}/document/delete/{doc}/{uid}", deleteDocument)

	log.Fatal(http.ListenAndServe(":5555", r))
}

func deleteCollection(params mamba.Dict) {
	collectionName := params["collectionName"]

	os.Remove("db/"+collectionName+".properties")
}