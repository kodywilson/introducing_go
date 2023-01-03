package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Name struct {
	Fname []string
}

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("hello.html"))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	name := Name{
		Fname: []string{"Bob", "John", "Tom"},
	}
	err := tpl.Execute(w, name)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8099", nil)
}
