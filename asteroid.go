package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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

	// UPDATE operation
	r.HandleFunc("/api/collection/{name}/document/update/{doc}", updateDocument)

	// DELETE operation
	r.HandleFunc("/api/collection/{name}/document/delete/{doc}", deleteDocument)

	log.Fatal(http.ListenAndServe(":5555", r))

}
