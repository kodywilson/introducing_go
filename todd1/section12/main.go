package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person // anonymous field (embedded) - no need to specify type
	ltk    bool
}

func main() {
	bob := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	chuck := person{
		first: "Chunk",
		last:  "Norris",
		age:   72,
	}
	fmt.Println(bob, chuck)
	fmt.Println(bob.first, bob.last)

	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa)
	// don't need embedded type name, but can use to
	// prevent name collisions
	fmt.Println(sa.first, sa.last, sa.person.age)

	creature := struct {
		name string
	}{
		name: "Bear",
	}
	fmt.Println(creature.name)
}
