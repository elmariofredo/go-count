package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogs() {

	if os.Getenv("LOG_FORMAT") == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

}

func main() {

	initLogs()

	log.Info("Starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/count", count)

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
