package main

import (
	"encoding/json"
	"fmt"
)

// ex1
type user struct {
	First string
	Age   int
}

// ex2
type User struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	// ex1 - marshall to json
	u1 := user{First: "Bob", Age: 23}
	u2 := user{First: "James", Age: 32}
	u3 := user{First: "Jill", Age: 65}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))

	// ex2
	jayceon := `[{"First":"Bob","Age":23},{"First":"James","Age":32},{"First":"Jill","Age":65}]`

	var Users []User // create slice of User
	// check error (there is no assignment on left!)
	// json.Unmarshall(slice of byte, pointer to slice)
	erro := json.Unmarshal([]byte(jayceon), &Users)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(Users)
}
