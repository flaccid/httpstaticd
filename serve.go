package httpstaticd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

var (
	code int = http.StatusTemporaryRedirect
)

func Serve(directory string, listenPort int, listings bool, accessLog bool) {
	log.Info("initialize httpstaticd")
	log.Debug("debug logging enabled")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/", healthCheckHandler)

	log.WithFields(log.Fields{
		"dir":  directory,
		"port": listenPort,
	}).Info("listening for requests")

	if listings {
		log.Info("directory listings enabled")
		http.Handle("/", handlers.LoggingHandler(os.Stdout, http.FileServer(http.Dir(directory))))

		if err := http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), nil); err != nil {
			log.Fatalf("fatality: %v", err)
		}
	} else {
		mux := http.NewServeMux()
		fileServer := http.FileServer(neuteredFileSystem{http.Dir(directory)})
		mux.Handle("/", http.StripPrefix("/", handlers.LoggingHandler(os.Stdout, fileServer)))

		if err := http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), mux); err != nil {
			log.Fatalf("fatality: %v", err)
		}
	}
}
