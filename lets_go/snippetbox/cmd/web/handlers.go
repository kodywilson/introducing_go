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

	app.render(w, http.StatusOK, "create.gohtml", data)
}

// snippetCreate handler function
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
	}

  // Get the values with the matching keys (title, content, etc.)
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// form data is string, but we want int for expires
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
			app.clientError(w, http.StatusBadRequest)
			return
	}

	// validate data
	fieldErrors := make(map[string]string)

	// check title length
	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot be blank"
  } else if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field cannot be more than 100 characters long"
  }

	// check content field
	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field cannot be blank"
  }

	// check expires
	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
  }

	// return errors if any are found
	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

  // write the snippet data to the DB
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
			app.serverError(w, err)
			return
	}

	// redirect to snippet page
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
