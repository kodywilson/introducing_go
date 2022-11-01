package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func todo(w http.ResponseWriter, r *http.Request) {

	data := PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Like this video", Done: false},
			{Item: "Reserve campsites", Done: false},
		},
	}

	err := tpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))

	mux.HandleFunc("/todo", todo)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":9091", mux))
}
