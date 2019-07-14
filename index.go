package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	users:=map[string]int {
		"a": 20,
		"b": 30,
		"c": 40,
		"d": 50,
	}

	router:=mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "signup.html")
	})
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})

	router.HandleFunc("/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		// name:=r.URL.Path[len("/users/"):]
		vars:=mux.Vars(r)
		name:=vars["name"]
		age:=users[name]
		fmt.Fprintf(w, "User name is %s %d years old", name, age)
	}).Methods("GET")

	http.ListenAndServe(":3001", router)

}

