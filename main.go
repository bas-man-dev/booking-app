package main

import "fmt"

func main() {

	var conferenceName = "Fantastic Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("\nWelcome to the %s booking application:\n", conferenceName)
	fmt.Printf("We have %d tickets and %d available\n", conferenceTickets, remainingTickets)
	fmt.Println("Continue here to get your tickets to attend the conference")

}
