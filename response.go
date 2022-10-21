package httpstaticd

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter, status int, contentType string, body string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", contentType)
	io.WriteString(w, body)
}
