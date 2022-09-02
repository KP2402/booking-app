package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("_______________________________________________________")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v seats available\n", conferenceTickets)
	fmt.Printf("Only %v seats remaining\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("_______________________________________________________")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask the user for the name
	fmt.Print("Enter your firstName: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter no. of tickets to buy: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %v seats remaining\n", remainingTickets)
	fmt.Println("_______________________________________________________")

}
