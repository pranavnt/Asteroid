package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(string(ioutil.ReadAll(r.Body))), &data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}

func signUp(w http.ResponseWriter, r *http.Request) {}

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

	if err != nil {
		return false
	}

	return true
}
