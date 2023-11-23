package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets int  = 50
	var remainingTickets uint = 50
	//array
	//var bookings[50]string
	//slice
   bookings := []string{}

	fmt.Printf("Welcome to %v application\n", conferenceName)
	fmt.Println("We have total of",conferenceTickets, "tickets and", remainingTickets, "are still available" )
	fmt.Println("Get your tickets here")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name 
		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your Email:")
		fmt.Scan(&email)
	
		fmt.Println("Enter your Tickets:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets{
			fmt.Printf("we have only %v tickets remaining , so you can't book %v tickets.\n", remainingTickets,userTickets)
			continue
		}
	
		if userTickets <= remainingTickets{
			remainingTickets = remainingTickets - userTickets 
			//array
			//bookings[0] = firstName +""+lastName
			//slice
			bookings=append(bookings,firstName +""+lastName)
		
			fmt.Printf("The Whole Array value: %v\n" , bookings)
			fmt.Printf("The First Array value: %v\n" , bookings[0])
			fmt.Printf("The Array type: %T\n" , bookings)
			fmt.Printf("The Array length: %v\n" , len(bookings))
		
			fmt.Printf("The Whole slice value: %v\n" , bookings)
			fmt.Printf("The First slice value: %v\n" , bookings[0])
			fmt.Printf("The slice type: %T\n" , bookings)
			fmt.Printf("The slice length: %v\n" , len(bookings))
	
			noTicketsBooked := remainingTickets == 0
	
			if noTicketsBooked{
				fmt.Printf("Our tickets are booked out.Thank you , come back later")
				break
			}
		}
		
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an confirmation email at %v Soon.\n" , firstName,lastName,userTickets,email)
		fmt.Printf("%v tickets are available for %v.\n", remainingTickets, conferenceName)
	}
	

}