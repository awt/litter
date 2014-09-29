package public

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/awt/litter/store"
)

type Message struct {
	Body string
	Signature string
	From string
}

type ApiHandler struct {
}


// Check friends' leets periodically and store them
// in our leet table


func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// receive direct message
	// receive notification of new leet
	// receive notification of mention

	var responseBody string
	code := 200

	if req.Header.Get("Accept") == "application/json" {
		requestBody := make([]byte, req.ContentLength)
		req.Body.Read(requestBody)

		// Build the response and send it
		responseBody, code = route(path, method, requestBody)
	} else {
		responseBody, code = route(path, method)
	}
	w.WriteHeader(code);
	fmt.Fprintf(w, responseBody)
}

func route(path string, method string, args ...interface{}) (body string, code int){

	if path == "/notify" && method == "POST" {
		// receive leet notificaiton

		var msg Message
		var requestBody = args[0].([]byte)
		json.Unmarshal(requestBody, &msg)
		store.CreateLeet(msg.Body)
		body = ""
		code = 200
	} else if path == "/leets" && method == "GET" {
		// send all leets before cutoff

		body = fmt.Sprintf("Fetching leets %s", "foo")
		code = 200
	} else if path == "/pubkey" && method == "GET" {
		// send gnupg public key

		body = fmt.Sprintf("Fetching leets %s", "foo")
		code = 200
	} else {
		body = ""
		code = 404
	}
	return body, code
}

