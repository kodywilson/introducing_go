package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// parse one file into tpl
	tpl, err := template.ParseFiles("one.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	// show contents of tpl
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// parse two more files
	tpl, err = tpl.ParseFiles("two.gmao", "vespa.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	// show contents of tpl
	// will only show first template added
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// execute each template individually
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
}
