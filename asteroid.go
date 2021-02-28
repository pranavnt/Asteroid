package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)
func dictToJson(dict map[string]interface{}) string {
	var hold = ""
	hold+="{"
    for a := 0; a < len(dict)-1; a++ {
		hold+=fmt.Sprintf("\"%s\":\"%s\",",//insert key here,//insert value here"ehhy")
	}
	hold+=fmt.Sprintf("\"%s\":\"%s\"",//insert key here,//insert value here"ehhy")
	hold+="}"
	fmt.Println(hold)
	return hold
}

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

	// UPDATE operation
	r.HandleFunc("/api/collection/{name}/document/update/{doc}", updateDocument)

	// DELETE operation
	r.HandleFunc("/api/collection/{name}/document/delete/{doc}", deleteDocument)
	
	log.Fatal(http.ListenAndServe(":5555", r))
	


	
}
