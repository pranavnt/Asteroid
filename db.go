package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func createCollection(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	os.Create("db/collections/" + name + ".properties")
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	collectionName := mux.Vars(r)["name"]
	filePath := "db/collections/" + collectionName + ".properties"

	fmt.Println(filePath)

	var req map[string]interface{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(filePath)

	json.Unmarshal([]byte(string(bytes)), &req)

	id, _ := gonanoid.New()
	uid := req["uid"]

	if isValidUid(uid.(string)) {
		req["_id"] = id

	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "data")
}

func isValidUid(uid string) bool {
	fmt.Println(uid)
	return true
}

func addDocument(key string, value string) {
	return
}
