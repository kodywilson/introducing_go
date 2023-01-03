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

	menu1 := menu{
		Breakfast: []string{"Bacon", "Cereal", "Pancakes"},
		Lunch:     []string{"Reuben", "Hamburger", "Salad"},
		Dinner:    []string{"Steak", "Lobster", "Fancy Soup"},
	}

	menu2 := menu{
		Breakfast: []string{"Bacon", "Cereal", "Pancakes"},
		Lunch:     []string{"Reuben", "Hamburger", "Salad"},
		Dinner:    []string{"Steak", "Lobster", "Fancy Soup"},
	}

	menu3 := menu{
		Breakfast: []string{"Bacon", "Cereal", "Pancakes"},
		Lunch:     []string{"Reuben", "Hamburger", "Salad"},
		Dinner:    []string{"Steak", "Lobster", "Fancy Soup"},
	}

	rest_menus := []menu{menu1, menu2, menu3}

	err := tpl.Execute(os.Stdout, rest_menus)
	if err != nil {
		log.Fatalln(err)
	}
}
