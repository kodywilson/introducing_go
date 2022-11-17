package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// later auto generate from build
const version = "1.0.0"

// configuration struct
type config struct {
	port int
	env  string
}

// application struct
type application struct {
	config config
	logger *log.Logger
}

// 11-16-2022
func main() {
	// declare cfg, config struct
	var cfg config

	// read command line arguments or supply defaults
	flag.IntVar(&cfg.port, "port", 7000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// declare app of type application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare HTTP server and set reasonable defaults
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
