package main

import (
	"fmt"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println(("Get your tickets here to attend"))

	var username string
	var userTickets int
	var totalUsers int = 0

	var bookings [50]string

	//  Pointer --> &
	fmt.Println("Enter your name:-")
	fmt.Scan(&username)
	fmt.Println("Enter the tickets you want to book:-")
	fmt.Scan(&userTickets)

	bookings[totalUsers] = username
	totalUsers = totalUsers + 1

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets\n", username, userTickets)
	fmt.Printf("We have %v tickets left\n", remainingTickets)
	fmt.Printf("length of array:- %v", len(bookings))
}
