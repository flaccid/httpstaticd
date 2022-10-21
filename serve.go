package httpstaticd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var (
	code    int = http.StatusTemporaryRedirect
	handler http.Handler
)

func Serve(directory string, listings bool, listenPort int, enableCors bool, accessLog bool) {
	log.Info("initialize httpstaticd")
	log.WithFields(log.Fields{
		"dir":  directory,
		"port": listenPort,
	}).Info("listening for requests")

	if listings {
		// TODO: support dir listing when index file exists
		log.Info("directory listings enabled")
		http.Handle("/", handlers.LoggingHandler(os.Stdout, http.FileServer(http.Dir(directory))))
		http.HandleFunc("/health", healthCheckHandler)
		handler = nil
	} else {
		fileServer := http.FileServer(neuteredFileSystem{http.Dir(directory)})
		mux := http.NewServeMux()
		mux.Handle("/", http.StripPrefix("/", handlers.LoggingHandler(os.Stdout, fileServer)))
		mux.HandleFunc("/health", healthCheckHandler)
		if enableCors {
			log.Info("cors support enabled")
			handler = cors.Default().Handler(mux)
		} else {
			handler = mux
		}
	}

	if err := http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), handler); err != nil {
		log.Fatalf("fatality: %v", err)
	}
}
