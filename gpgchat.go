package main

import (
	"encoding/json"
	"net/http"
	//"html"
	"fmt"
	"log"
	//"os/exec"
)
// Start tor hidden service

// Register and update namecoin address

// Start http server

type Message struct {
	Body string
	Signature string
	From string

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var msg Message
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		json.Unmarshal(body, &msg)
		fmt.Fprintf(w, "%s", msg.Signature)
	})
	log.Fatal(http.ListenAndServe(":7777", nil))
}
