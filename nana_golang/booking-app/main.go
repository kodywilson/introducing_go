package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for their name
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %s booked %d tickets.\n", userName, userTickets)

}
