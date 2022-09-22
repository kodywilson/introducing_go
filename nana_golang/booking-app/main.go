package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %s booked %d tickets.\n", userName, userTickets)

}
