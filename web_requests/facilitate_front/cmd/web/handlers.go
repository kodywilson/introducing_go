package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"github.com/kodywilson/snippetbox/internal/models"
)

// store snippet form info
type snippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}

// Home handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
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
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
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

// snippetCreate returns form for creating new snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
  data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.gohtml", data)
}

// snippetCreatePost POST new snippet
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
	}

	// form data is string, but we want int for expires
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
	}

	// Allows us to store form data in case user has to submit again
	form := snippetCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     expires,
		FieldErrors: map[string]string{},
  }

	// check title length
	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
  } else if utf8.RuneCountInString(form.Title) > 100 {
		form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
  }

	// check content field
	if strings.TrimSpace(form.Content) == "" {
		form.FieldErrors["content"] = "This field cannot be blank"
  }

	// check expires
	if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
		form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
  }

	// return errors if any are found
	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.gohtml", data)
		return
	}

  // write the snippet data to the DB
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
			app.serverError(w, err)
			return
	}

	// redirect to snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
