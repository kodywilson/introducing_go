package main

import (
	"fmt"
)

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Tex"] = "tex@gmail.com"
	emails["Ged"] = "ged@gmail.com"

	fmt.Println(emails)
	delete(emails, "Ged")
	fmt.Println(emails)

	books := map[string]int {
		"Chamber of secrets": 256,
		"Lord of the Rings": 735,
	}
	fmt.Println(books)
}