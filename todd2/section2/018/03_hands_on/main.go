package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		{
			Name:    "California",
			Address: "100 Fake Drive",
			City:    "Los Angeles",
			Zip:     55454,
			Region:  "Southern",
		},
		{
			Name:    "Motel 6",
			Address: "106 Sweet Hills Drive",
			City:    "San Francisco",
			Zip:     53987,
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
