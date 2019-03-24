package httpstaticd

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, body)
}
