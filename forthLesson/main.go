package main

import (
	"fmt"
	"net/http"
)

// should use templates, however this is just a simple example of how it can also be done
func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1> Hello there </h1>")
	fmt.Fprintf(w, "<p> My name is <strong>%s</strong> and i am %s </p>", "Mustafa", "a software engineer")
}

// should never do this
func badindexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `
		<h1> Hello there </h1>
		<p> My name is <strong>%s</strong> and i am %s </p>
		`, "mustafa", "a software engineer")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bad", badindexHandler)
	http.ListenAndServe(":8000", nil)
}
