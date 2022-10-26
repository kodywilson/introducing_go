package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	name := os.Args[1]

	str := fmt.Sprintf(`
  <!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	// create a file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close() // make sure file closes before program exits

	// write str to file
	io.Copy(nf, strings.NewReader(str))

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
