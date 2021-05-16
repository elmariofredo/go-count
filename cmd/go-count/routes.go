package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.welcome)
	mux.HandleFunc("/count", app.count)

	return mux
}
