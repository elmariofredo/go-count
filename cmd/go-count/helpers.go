package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/go-redis/redis/v8"
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

	if os.Getenv("SERVER_PORT_NUMBER") != "" {
		app.config.serverPortNumber = os.Getenv("SERVER_PORT_NUMBER")
	}
	if os.Getenv("REDIS_PORT_NUMBER") != "" {
		app.config.redisPortNumber = os.Getenv("REDIS_PORT_NUMBER")
	}
	if os.Getenv("REDIS_HOSTNAME") != "" {
		app.config.redisHostname = os.Getenv("REDIS_HOSTNAME")
	}
	if os.Getenv("REDIS_PASSWORD") != "" {
		app.config.redisPassword = os.Getenv("REDIS_PASSWORD")
	}
	if os.Getenv("REDIS_DB") != "" {
		app.config.redisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	}

}

func (app *application) dbConnect() {

	app.ctx = context.Background()

	app.rdb = redis.NewClient(&redis.Options{
		Addr:     app.config.redisHostname + ":" + app.config.redisPortNumber,
		Password: app.config.redisPassword,
		DB:       app.config.redisDB,
	})

}

func (app *application) start() {

	log.Info("Starting")

	// Valerie first code
	VALERIEMJHYUI5RGVXS := "👩🏼‍🦰"
	log.Info(VALERIEMJHYUI5RGVXS)

	// Jasmine first code
	J6KNMKIYTREDSW324 := "🦄"
	log.Info(J6KNMKIYTREDSW324)

	err := http.ListenAndServe(app.config.serverAddress+":"+app.config.serverPortNumber, app.routes())

	log.Fatal(err)

}

func (app *application) serverError(w http.ResponseWriter, err error) {

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	log.Fatal(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)

}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
