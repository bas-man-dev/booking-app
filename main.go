package main

import (
	"fmt"
	"time"
)

func main() {

	conferenceName := "Fantastic Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("\nWelcome to the %v booking application:\n", conferenceName)
	fmt.Printf("We have %v tickets and %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Continue here to get your tickets to attend the conference")

	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like? ")
	fmt.Scan(&userTickets)

	now := time.Now()

	remainingTickets -= userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you for the purchase %v %v.\nYou bought %v tickets.\nYou will receive a confirmation email to:\n%v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Purchased at: %v:%v on the %v %v %v \n", now.Hour(), now.Minute(), now.Day(), now.Month(), now.Year())

	fmt.Println(remainingTickets)
	fmt.Printf("Everything: %v\n", bookings)
	fmt.Printf("First element: %v\n", bookings[0])
	fmt.Printf("Element type: %T\n", bookings)
	fmt.Printf("Size : %v\n", len(bookings))
}
