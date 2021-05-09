package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	portNumber string = "3000"
)

func setup() {

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

	if os.Getenv("PORT_NUMBER") != "" {
		portNumber = os.Getenv("PORT_NUMBER")
	}

}

func main() {

	setup()

	log.Info("Starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/count", count)

	err := http.ListenAndServe(":"+portNumber, mux)

	log.Fatal(err)
}
