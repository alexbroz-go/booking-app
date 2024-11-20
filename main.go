package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v of tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	userName = "Tom"
	userTickets = 10
	fmt.Printf("User %v booked %v tickets", userName, userTickets)

}
