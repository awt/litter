package private

import (
	"net/http"
	"fmt"
	"github.com/awt/litter/store"
)

type ApiHandler struct {
}

func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	body, code := route(path, method)
	w.WriteHeader(code);
	fmt.Fprintf(w, body)
}

func route(path string, method string) (body string, code int){

	if path == "/" && method == "POST" {
		// publish leet
		store.Exec()
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
