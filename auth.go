package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/magiconair/properties"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(string(bytes)), &data)

	guess := data["password"].(string)

	p := properties.MustLoadFile("db/users.properties", properties.UTF8)

	val, key := p.Get(data["username"].(string))

	if !key {
		fmt.Fprintf(w, "User not registered")
	}

	bytes = []byte(val)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(string(bytes)), &data)

	if checkPassword(data["password"].(string), guess) {
		fmt.Fprintf(w, data["userID"].(string))
	} else {
		fmt.Fprint(w, "Incorrect password")
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal([]byte(string(bytes)), &data)

	if erro != nil {
		log.Fatal(err)
	}

	fmt.Println("Username: " + data["username"].(string))

	fmt.Fprintf(w, addUser(data["username"].(string), data["password"].(string)))
}

func addUser(username string, password string) (id string) {
	password = hashPassword(password)
	p := properties.MustLoadFile("db/users.properties", properties.UTF8)

	id, _ = gonanoid.New()
	fmt.Println("Salt hashed password: " + password)

	fmt.Println("User ID: " + id)

	val := "{" + "\"username\": \"" + username + "\" ,\"password\": \"" + password + "\", \"userID\": \"" + id + "\"}"

	p.Set(username, val)

	err := ioutil.WriteFile("db/users.properties", []byte(p.String()), 0777)

	if err != nil {
		fmt.Println(err)
	}

	return
}

func hashPassword(password string) (hashedPassword string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	hashedPassword = string(hash)
	return
}

func checkPassword(hashedPassword string, guess string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(guess))

	return err == nil
}
