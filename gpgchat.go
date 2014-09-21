package main

import (
	//"encoding/json"
	"net/http"
	"html"
	"fmt"
	"log"
)
// Start tor hidden service

// Register and update namecoin address

// Start http server

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":7777", nil))
}
