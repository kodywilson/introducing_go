package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/kodywilson/snippetbox/internal/models"
)

// Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// make sure "/" was sent
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// generate list of most recent snippets
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// initialize data struct and add snippets
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// display data on home page
	app.render(w, http.StatusOK, "home.gohtml", data)
}

// snippetView handler function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// extract id from query string and test if integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// GET data for snippet based on ID
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Initialize data struct and add snippet
	data := app.newTemplateData(r)
	data.Snippet = snippet

	// display selected snippet
	app.render(w, http.StatusOK, "view.gohtml", data)
}

// snippetCreate handler function
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// make sure request is a POST request
	if r.Method != http.MethodPost {
		// send 405, "Method Not Allowed", and let user know allowed verbs
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// first insert some dummy values
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	// create snippet
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect to snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
