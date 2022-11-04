package main

import (
	"fmt"
	"net/http"
)

// POST new movie
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// GET existing movie by ID
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Show movie
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
