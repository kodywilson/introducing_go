package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type parson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// marshall
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Money",
		Last:  "Penny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Money","Last":"Penny","Age":27}]`
	bes := []byte(s)

	// slice of parsons to store my unmarshalled json
	var peeps []parson
	erro := json.Unmarshal(bes, &peeps)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(peeps)
}
