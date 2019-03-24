package httpstaticd

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var (
	code int = http.StatusTemporaryRedirect
)

func Serve(directory string, listenPort int, listings bool) {
	log.Info("initialize httpstaticd")
	log.Debug("debug logging enabled")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/", healthCheckHandler)

	// work out how to normalize
	if listings {
		log.Info("directory listings enabled")
		http.Handle("/", http.FileServer(http.Dir(directory)))
		log.WithFields(log.Fields{
			"dir":  directory,
			"port": listenPort,
		}).Info("listening for requests")
		if err := http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), nil); err != nil {
			log.Fatalf("fatality: %v", err)
		}
	} else {
		mux := http.NewServeMux()
		fileServer := http.FileServer(neuteredFileSystem{http.Dir(directory)})
		mux.Handle("/", http.StripPrefix("/", fileServer))
		log.WithFields(log.Fields{
			"dir":  directory,
			"port": listenPort,
		}).Info("listening for requests")
		if err := http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), mux); err != nil {
			log.Fatalf("fatality: %v", err)
		}
	}
}
