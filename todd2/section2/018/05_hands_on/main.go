package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	daily_menu := menu{
		Breakfast: []string{"Bacon", "Cereal", "Pancakes"},
		Lunch:     []string{"Reuben", "Hamburger", "Salad"},
		Dinner:    []string{"Steak", "Lobster", "Fancy Soup"},
	}

	err := tpl.Execute(os.Stdout, daily_menu)
	if err != nil {
		log.Fatalln(err)
	}
}
