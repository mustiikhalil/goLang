package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-http-utils/logger"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello there </h1>")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", logger.Handler(router, os.Stdout, logger.DevLoggerType))
}
