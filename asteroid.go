package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func makeCollection(w http.ResponseWriter, r *http.Request) {

}

func makeDocument(w http.ResponseWriter, r *http.Request) {

}

func getCollection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := r.URL.Query()
	fmt.Println(vars)
	fmt.Println(params)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getDocument(w http.ResponseWriter, r *http.Request) {

}

func updateDocument(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", login)
	r.HandleFunc("/signUp", signUp)

	//write operations
	r.HandleFunc("/api/collection/{name}", createCollection).Methods("POST")
	r.HandleFunc("/api/collection/{name}/document", makeDocument).Methods("POST")

	//read operations
	r.HandleFunc("/api/collection/{name}", getCollection).Methods("GET")
	r.HandleFunc("/api/collection/{name}/document/{doc}", getDocument).Methods("GET")

	//update operation
	r.HandleFunc("/api/collection/{name}/document/{doc}", updateDocument).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5555", r))

}
