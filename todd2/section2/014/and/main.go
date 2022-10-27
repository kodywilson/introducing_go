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

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "Joe",
		Motto: "Mama",
		Admin: true,
	}

	u2 := user{
		Name:  "George",
		Motto: "I'd rather be lucky than good",
		Admin: true,
	}

	u3 := user{
		Name:  "Bob",
		Motto: "I like apples",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
