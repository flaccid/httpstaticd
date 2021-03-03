package httpstaticd

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func logRequest(r *http.Request) {
	// we only log in debug mode due to possible exposure of PI data in request uri
	log.WithFields(log.Fields{
		"method": r.Method,
	}).Debug(r.URL.String())
}
