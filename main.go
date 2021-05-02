package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	println("Starting")

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Ahoj"))
	})

	counter := 1

	http.HandleFunc("/count", func(rw http.ResponseWriter, r *http.Request) {
		
		counter++

		rw.Write([]byte(strconv.Itoa(counter)))

	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
