package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.welcome)
	mux.HandleFunc("/count", app.count)
	mux.HandleFunc("/ping", app.ping)

	return mux
}
