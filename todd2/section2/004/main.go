package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close() // make sure file closes before program exits

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
