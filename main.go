package main

import (
	"log"
	"net/http"
	"strconv"
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
}

func count(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {

		counter = 0

		rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

		return

	}

	counter++

	rw.Write([]byte(`{"count": ` + strconv.Itoa(counter) + `}`))

}

func main() {
	println("Starting")

	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)
	mux.HandleFunc("/count", count)

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
