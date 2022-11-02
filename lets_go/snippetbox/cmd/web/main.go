package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	// set up config via command line arguments or use defaults
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "Web server port")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// Initialize a new servemux
	mux := http.NewServeMux()

	// create file server for statics
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	// Register handlers
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)     // register snippet view handler
	mux.HandleFunc("/snippet/create", snippetCreate) // register snippet create handler

	// Start web server
	log.Printf("Starting server on %s\n", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
