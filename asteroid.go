package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", login)
	r.HandleFunc("/signUp", signUp)
	http.Handle("/", r)

	fmt.Println("hello world")
}
