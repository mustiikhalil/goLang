package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "about page")
}

func main() {
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
