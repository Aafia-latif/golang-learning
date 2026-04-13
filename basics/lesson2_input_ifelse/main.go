package main

import "fmt"

func main() {
	// Variables for user input
	var name string
	var age int
	var ticketPrice int

	// Ask for user input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// Validation
	if age <= 0 {
		fmt.Print("Invalid age entered ❌")
		return
	}

	fmt.Println("\n----- Ticket Details -----")

	//Decision making with if-else
	if age >= 18 {
		ticketPrice = 1000
		fmt.Printf("Hello %s , you are an Adult ✅\n", name)
		fmt.Println("Ticket Type: Adult")
	} else if age >= 13 {
		ticketPrice = 700
		fmt.Printf("Hello %s , you are a Teen ⚠️\n", name)
		fmt.Println("Ticket Type: Teen")
		fmt.Println("Parental guidance is recommended.")
	} else {
		ticketPrice = 500
		fmt.Printf("Hello %s , you are a Child 👶\n", name)
		fmt.Println("Ticket Type: Child")
		fmt.Println("Parental guidance is required.")
	}

	fmt.Println("------------------------------")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Ticket Price:", ticketPrice)
	fmt.Println("------------------------------")
}
