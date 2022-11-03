package main

import (
	"fmt"
	"net/http"
)

// declare handler which writes plain-text response with the following
// information: application status, operating environment, and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
