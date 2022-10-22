package main

import "fmt"

type Player struct {
	Name string
	Str int
	Father *Player
}

type Mob struct {
	Name string
	Int int
}

func main() {
	bob := Player{
		Name: "Robert",
		Str: 16,
	}
	fmt.Println(bob.Name)

	darrel := &Player{Name: "Darrel"}
	darrel.Str = 12
	fmt.Println(darrel.Name, darrel.Str)

	Super(darrel)
	fmt.Println(darrel.Name, darrel.Str)

	wiz := &Mob{"Fizban", 14}
	wiz.Souper()
	fmt.Println(wiz.Int) // will show two added from Souper() = 18

	bob = NewPlayer("Bob", 16)
	fmt.Println(bob.Name)

	george := &Player{
		Name: "George",
		Str: 14,
		Father: &Player {
			Name: "Johny",
			Str: 16,
			Father: nil,
		},
	}

  // george is a pointer so * to dereference
	fmt.Println(*george.Father)
}

func Super(p *Player) {
	p.Str += 2
}

func (m *Mob) Souper() {
	m.Int += 2
}

func NewPlayer(name string, str int) Player {
	return Player{
		Name: name,
		Str: str,
	}
}
