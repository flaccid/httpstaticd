// NOTE: not yet used/implemented

package httpstaticd

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func logRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
	}).Debug(r.URL.String())
}
