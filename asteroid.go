package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
)

func makeCollection(w http.ResponseWriter, r *http.Request) {

}

func makeDocument(w http.ResponseWriter, r *http.Request) {

}

func getCollection(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	// vars := mux.Vars(r)
	params := r.URL.Query()
	// collection := vars["name"]
	usrId := params["userID"][0]

	p := properties.MustLoadFile("db/users.properties", properties.UTF8)
	keys := p.Keys()
	fmt.Println(keys)

	for a := 0; a < len(keys); a++ {
		val, key := p.Get(keys[a])

		if key == false {
			fmt.Fprintf(w, "User not registered")
		}

		bytes := []byte(val)
		json.Unmarshal([]byte(string(bytes)), &data)

		fmt.Println(data["userID"])
		if data["userID"] == usrId {
			fmt.Fprintf(w, "Access Granted")
		}
	}

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
	r.HandleFunc("/api/collection/{name}/document", createDocument).Methods("POST")

	//read operations
	r.HandleFunc("/api/collection/{name}", getCollection).Methods("GET")
	r.HandleFunc("/api/collection/{name}/document/{doc}", getDocument).Methods("GET")

	//update operation
	r.HandleFunc("/api/collection/{name}/document/{doc}", updateDocument).Methods("PUT")

	//signup
	r.HandleFunc("/signup", signUp).Methods("POST")

	log.Fatal(http.ListenAndServe(":5555", r))

}
