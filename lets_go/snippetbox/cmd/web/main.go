package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servemux
	mux := http.NewServeMux()

	// create file server for statics
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register handlers
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)     // register snippet view handler
	mux.HandleFunc("/snippet/create", snippetCreate) // register snippet create handler

	// Start web server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
