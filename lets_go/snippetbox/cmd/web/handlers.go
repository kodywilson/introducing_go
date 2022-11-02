package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// make sure "/" was sent
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// extract id from query string and test if integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// make sure request is a POST request
	if r.Method != http.MethodPost {
		// send 405, "Method Not Allowed", and let user know allowed verbs
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
