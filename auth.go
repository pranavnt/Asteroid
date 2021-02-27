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

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
	}

	erro := json.Unmarshal([]byte(string(bytes)), &data)

	if erro != nil {
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
