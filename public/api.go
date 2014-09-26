package public

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Message struct {
	Body string
	Signature string
	From string
}

type ApiHandler struct {
}


func (h *ApiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
		// receive notification of new leet
		// receive direct message
		// receive notification of mention

		var msg Message
		body := make([]byte, req.ContentLength)
		req.Body.Read(body)
		json.Unmarshal(body, &msg)
		fmt.Fprintf(w, "%s", msg.Signature)
}



