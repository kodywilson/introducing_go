package main

import (
	"log"
	"os"
	"text/template"
)

// don't forget this is a pointer to template.Template!
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
