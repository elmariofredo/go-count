package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func (app *application) setup() {

	log.Info("Preparing")

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

func (app *application) start() {

	log.Info("Starting")

	err := http.ListenAndServe(":"+portNumber, app.routes())

	log.Fatal(err)

}

// func (app *application) serverError(w http.ResponseWriter, err error) {

// 	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
// 	log.Fatal(trace)

// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

// }

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)

}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
