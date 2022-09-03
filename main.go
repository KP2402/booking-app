package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets //unsigned integer since number of tickets cannot be negative

	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("_______________________________________________________")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v seats available\n", conferenceTickets)
	fmt.Printf("Only %v seats remaining\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("_______________________________________________________")

	// Arrays and slices
	// Arrays have fixed size. Slices has dynamic size
	// var bookings = [50]string{"Prashant", "Tom"}
	// var bookings [50]string // alternative initalization of above
	// bookings[0] = "Prashant"
	// bookings[1] = "Tom"
	var bookings = [50]string{} //Arrays
	var bookingEmails []string  //Slices

	var firstName string
	var lastName string
	var email string
	var userTickets uint //unsigned integer since number of tickets cannot be negative

	// ask the user for the name
	fmt.Print("Enter your firstName: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter no. of tickets to buy: ")
	fmt.Scan(&userTickets)

	bookings[0] = firstName + " " + lastName
	bookingEmails = append(bookingEmails, email)
	remainingTickets = remainingTickets - userTickets

	fmt.Println("_______________________________________________________")
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value in array: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array Size: %v\n", len(bookings))
	fmt.Println("_______________________________________________________")

	fmt.Printf("The whole slice: %v\n", bookingEmails)
	fmt.Printf("The first value in slice: %v\n", bookingEmails[0])
	fmt.Printf("Slice Type: %T\n", bookingEmails)
	fmt.Printf("Slice Size: %v\n", len(bookingEmails))
	fmt.Println("_______________________________________________________")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get the confirmation on your email %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Only %v seats remaining\n", remainingTickets)
	fmt.Println("_______________________________________________________")

}
