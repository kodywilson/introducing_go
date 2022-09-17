package main

import "fmt"

type Player struct {
	Name string
	Str int
}

func main() {
	bob := Player{
		Name: "Bob",
		Str: 16,
	}
	fmt.Println(bob.Name)

	darrel := &Player{Name: "Darrel"}
	darrel.Str = 12
	fmt.Println(darrel.Name, darrel.Str)

	Super(darrel)
	fmt.Println(darrel.Name, darrel.Str)
}

func Super(p *Player) {
	p.Str += 2
}