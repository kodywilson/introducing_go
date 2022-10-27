package main

import (
	"log"
	"os"
	"text/template"
)

// don't forget this is a pointer to template.Template!
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four"}

	data := struct {
		Words []string
	}{
		xs,
	}

	// err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
