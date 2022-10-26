package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// parse one file into tpl
	tpl, err := template.ParseGlob("templates/*.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	// show contents of tpl (only shows first one added)
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute each template individually
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
