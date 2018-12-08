package main

import (
	"fmt"
	"net/http"

	"github.com/mustii/stringutil"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa, go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "about page")
}

func main() {
	fmt.Println(stringutil.Reverse("welcome"))
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
