package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// CREATE
func createCollection(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	params := r.URL.Query()
	usrId := params["userID"][0]
	if(hasAccess(usrId)){
		os.Create("db/collections/" + name + ".properties")
		json.NewEncoder(w).Encode(map[string]bool{"status": true})
	}
	
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

	if hasAccess(uid.(string)) {
		req["_id"] = id

	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "data")
}

func addDocument(collection string, key string, value string) {
	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)
	p.Set(key, value)
	return
}

// READ
// func readCollection(w http.ResponseWriter, r *http.Request) {
// 	// params := r.URL.Query()
// 	// usrId := params["userID"][0]
// 	vars := mux.Vars(r)
// 	collection := vars["name"]
	
// 	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)
// 	fmt.Println(pupd)

// 	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
// }

func readDocument(w http.ResponseWriter, r *http.Request) {
	// params := r.URL.Query()
	// usrId := params["userID"][0]
	vars := mux.Vars(r)
	collection := vars["name"]
	doc :=vars["doc"]
	
	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)
	val, key := p.Get(doc)
	if(key==false){
		fmt.Fprintf(w,"does not exist")
	}

	fmt.Println(val)
	json.NewEncoder(w).Encode(val)

}

// UPDATE
func updateDocument(w http.ResponseWriter, r *http.Request) {

}

// DELETE
func deleteDocument(w http.ResponseWriter, r *http.Request) {

}

//returns if user has access to said data
func hasAccess(usrID string) bool {
	var data map[string]interface{}

	p := properties.MustLoadFile("db/users.properties", properties.UTF8)
	keys := p.Keys()
	fmt.Println(keys)

	for a := 0; a < len(keys); a++ {
		val, key := p.Get(keys[a])

		if key == false {
			return false
		}

		bytes := []byte(val)
		json.Unmarshal([]byte(string(bytes)), &data)

		if data["userID"] == usrID {
			return true
		}
	}
	return false
}

// func convertJson() string {

// }