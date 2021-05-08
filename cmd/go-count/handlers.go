package main

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	counter int = 0
)

func welcome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Ahoj"))

	log.WithFields(log.Fields{
		"path": r.URL.Path,
	}).Debug("Ahoj")

}

func count(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {

		counter = 0

		rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

		log.WithFields(log.Fields{
			"path": r.URL.Path,
		}).Info("Rest")

		return

	}

	counter++

	rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

	log.WithFields(log.Fields{
		"path":  r.URL.Path,
		"count": strconv.Itoa(counter),
	}).Debug("counting")

}