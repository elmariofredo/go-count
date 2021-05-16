package main

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func (app *application) welcome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Ahoj"))

	log.WithFields(log.Fields{
		"path": r.URL.Path,
	}).Debug("Ahoj")

}

func (app *application) count(rw http.ResponseWriter, r *http.Request) {

	counterValue, err := app.rdb.Get(app.ctx, "counter").Result()
	if err != nil {
		counterValue = "0"
	}

	var counter, _ = strconv.Atoi(counterValue)

	rw.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {

		counter := 0

		setErr := app.rdb.Set(app.ctx, "counter", strconv.Itoa(counter), 0).Err()
		if setErr != nil {
			app.serverError(rw, setErr)
			return
		}

		rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

		log.WithFields(log.Fields{
			"path": r.URL.Path,
		}).Info("Rest")

		return

	}

	counter++

	setErr := app.rdb.Set(app.ctx, "counter", strconv.Itoa(counter), 0).Err()
	if setErr != nil {
		app.serverError(rw, setErr)
		return
	}

	rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

	log.WithFields(log.Fields{
		"path":  r.URL.Path,
		"count": strconv.Itoa(counter),
	}).Debug("counting")

}
