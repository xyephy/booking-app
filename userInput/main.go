package main

import "fmt"

func main() {
	// variable definition
	conferenceName := "Go Conference" //short variable definition
	var remainingTickets uint = 50
	// constant definition
	const conferenceTickets int = 50

	fmt.Printf("conferenceTickets is %T, remainingTckets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// adding data types to variables
	var userName string //its of string data type
	var userTickets int // its of integer data type
	// ask user for their name
	fmt.Scan(userName)

	// ask user for number of tickets
	userTickets = 4
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
