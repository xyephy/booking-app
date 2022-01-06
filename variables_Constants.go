package main

import "fmt"

func main() {
	// variable definition
	var conferenceName = "Go Conference"
	var remainingTickets = 50
	// constant definition
	const conferenceTickets = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
