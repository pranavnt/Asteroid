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
	if hasAccess(usrId) {
		os.Create("db/collections/" + name + ".properties")
		json.NewEncoder(w).Encode(map[string]bool{"status": true})
	}

}

func createDocument(w http.ResponseWriter, r *http.Request) {
	collectionName := mux.Vars(r)["name"]
	filePath := "db/collections/" + collectionName + ".properties"

	var req map[string]interface{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(string(bytes)), &req)

	id, _ := gonanoid.New()
	uid := req["uid"]

	if hasAccess(uid.(string)) {
		p := properties.MustLoadFile(filePath, properties.UTF8)
		req["_id"] = id

		val := dictToJson(req)
		p.Set(id, val)

		ioutil.WriteFile(filePath, []byte(p.String()), 0777)
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Document created!!")

	fmt.Fprintf(w, id)
}

// READ
func readDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["name"]
	doc := vars["doc"]

	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)
	val, key := p.Get(doc)

	if key == false {
		fmt.Fprintf(w, "does not exist")
	}

	json.NewEncoder(w).Encode(json.RawMessage(val))
}

func readDocumentsById(w http.ResponseWriter, r *http.Request) {
	var documents []string
	vars := mux.Vars(r)
	collection := vars["name"]
	uid := vars["uid"]

	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)

	for _, key := range p.Keys() {
		val, _ := p.Get(key)

		var data map[string]interface{}

		json.Unmarshal([]byte(val), &data)

		if data["uid"].(string) == uid {
			documents = append(documents, dictToJson(data))
		}

	}

	respData := "{ \"data\": [" + strings.Join(documents, ",") + "]}"

	json.NewEncoder(w).Encode(json.RawMessage(respData))

}

func readAllDocuments(w http.ResponseWriter, r *http.Request) {
	var documents []string
	vars := mux.Vars(r)
	collection := vars["name"]

	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)

	for _, key := range p.Keys() {
		val, _ := p.Get(key)

		var data map[string]interface{}

		json.Unmarshal([]byte(val), &data)

		documents = append(documents, dictToJson(data))

	}

	respData := "{ \"data\": [" + strings.Join(documents, ",") + "]}"

	json.NewEncoder(w).Encode(json.RawMessage(respData))
}

// UPDATE
func updateDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["name"]
	doc := vars["doc"]

	var req map[string]interface{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(string(bytes)), &req)

	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)

	val, _ := p.Get(doc)

	var data map[string]interface{}

	json.Unmarshal([]byte(val), &data)

	if req["uid"] == data["uid"] {
		p.Set(doc, dictToJson(req))
	} else {
		fmt.Println("User does not have permission to modify this")
	}

	fmt.Println("Document updated!")

	ioutil.WriteFile(("db/collections/" + collection + ".properties"), []byte(p.String()), 0777)

	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

// DELETE
func deleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["name"]
	doc := vars["doc"]
	uid := vars["uid"]

	p := properties.MustLoadFile(("db/collections/" + collection + ".properties"), properties.UTF8)

	val, _ := p.Get(doc)

	var data map[string]interface{}

	json.Unmarshal([]byte(val), &data)

	if data["uid"] == uid {
		p.Delete(doc)
	}

	ioutil.WriteFile(("db/collections/" + collection + ".properties"), []byte(p.String()), 0777)

	fmt.Println("Document deleted!")
}

//returns if user has access to said data
func hasAccess(usrID string) bool {
	var data map[string]interface{}

	p := properties.MustLoadFile("db/users.properties", properties.UTF8)
	keys := p.Keys()

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

func dictToJson(dict map[string]interface{}) string {
	var hold = ""
	hold += "{"

	for k, v := range dict {
		hold += fmt.Sprintf("\"%s\":\"%s\",", k, v)
	}

	hold = strings.TrimSuffix(hold, ",")

	hold += "}"

	return hold
}
