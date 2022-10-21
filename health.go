package httpstaticd

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// we are always healthy!
	sendResponse(w, http.StatusOK, "application/json", `{"healthy": true}`)
}
