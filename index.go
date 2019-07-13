package main

import (
	"fmt"
	"net/http"
)

func main() {

	users:=map[string]int {
		"a": 20,
		"b": 30,
		"c": 40,
		"d": 50,
	}

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Saksit Jantaraplin")
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		name:=r.URL.Path[len("/users/"):]
		age:=users[name]
		fmt.Fprintf(w, "User name is %s %d years old", name, age)
	})
	http.ListenAndServe(":8080", nil)

}

