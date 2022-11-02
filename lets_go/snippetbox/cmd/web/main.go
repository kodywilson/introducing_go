package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	// set up loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize app struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Start web server
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s\n", cfg.addr)
	// call ListenAndServe() on srv http struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
