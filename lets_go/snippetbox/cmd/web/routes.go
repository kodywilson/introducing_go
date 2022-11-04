package main

import "net/http"

func (app *application) routes() http.Handler {
	// Initialize a new servemux
	mux := http.NewServeMux()

	// create file server for statics
	//fileServer := http.FileServer(http.Dir(cfg.staticDir))
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register handlers
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)     // register snippet view handler
	mux.HandleFunc("/snippet/create", app.snippetCreate) // register snippet create handler

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
