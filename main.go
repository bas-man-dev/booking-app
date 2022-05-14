package main

import (
	"fmt"
	"time"
)

func main() {

	var conferenceName = "Fantastic Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("\nWelcome to the %v booking application:\n", conferenceName)
	fmt.Printf("We have %v tickets and %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Continue here to get your tickets to attend the conference")

	var userName string
	var userTickets int

	fmt.Println("Please enter your name: ")
	fmt.Scan(&userName)
	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)

	now := time.Now()

	// userName = "Jon"
	// userTickets = 2

	fmt.Printf("User: %v || Tickets: %v \n", userName, userTickets)
	fmt.Println("Purchased at: ", now.Hour(), ":", now.Minute(), now.Day(), now.Month(), now.Year())

}
