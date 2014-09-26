package private

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/awt/litter/store"
	//"github.com/awt/litter/config"
)

type Message struct {
	Body string
}

type ApiHandler struct {
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// process the json in the request body

	var msg Message
	requestBody := make([]byte, req.ContentLength)
	req.Body.Read(requestBody)
	json.Unmarshal(requestBody, &msg)

	// Build the response and send it
	responseBody, code := route(path, method, msg)
	w.WriteHeader(code);
	fmt.Fprintf(w, responseBody)
}

func route(path string, method string, msg Message) (body string, code int){

	if path == "/" && method == "POST" {
		// publish leet
		store.Leet(msg.Body)
		body = ""
		code = 200
	} else if path == "/follow" && method == "PUT" {
		// follow litter name
		body = fmt.Sprintf("Following name %s", "foo")
		code = 200
	} else if path == "/" && method == "GET" {
		// get leets of followed names
		body = fmt.Sprintf("Fetching leets %s", "foo")
		code = 200
	} else {
		body = ""
		code = 404
	}
	return body, code
}
