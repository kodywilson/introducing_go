package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

// ex 4 and 5 - sort custom
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByFirst []user

func (l ByFirst) Len() int           { return len(l) }
func (l ByFirst) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByFirst) Less(i, j int) bool { return l[i].First < l[j].First }

func main() {
	// ex1 - marshall to json
	u1 := user{First: "Bob", Age: 32}
	u2 := user{First: "James", Age: 62}
	u3 := user{First: "Jill", Age: 55}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))

	// ex2 - unmarshall json to slice of structs
	jayceon := `[{"First":"Bob","Age":23},{"First":"James","Age":32},{"First":"Jill","Age":65}]`

	var Users []User // create slice of User
	// check error (there is no assignment on left!)
	// json.Unmarshall(slice of byte, pointer to slice)
	erro := json.Unmarshal([]byte(jayceon), &Users)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(Users)

	for i, user := range Users {
		fmt.Println("Person #", i)
		fmt.Println(user.First)
	}

	// ex3 - use json.NewEncoder
	json.NewEncoder(os.Stdout).Encode(users[0])

	// ex4 and 5 - sort for int and string
	// sort strings
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	// sort ints
	numbers := []int{10, 2, 5, 99, 232, 55, 3, 1, 100}
	sort.Ints(numbers)
	fmt.Println(numbers)

	// sort slice of structs, first by age, then by first
	sort.Sort(ByAge(users))
	fmt.Println(users)
	fmt.Println("------------------------------")
	sort.Sort(ByFirst(users))
	fmt.Println(users)

}
