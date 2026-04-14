package main

import "fmt"

func main() {
	var users int
	var totalRevenue int
	fmt.Print("Enter number of users:")
	fmt.Scan(&users)

	//Loop for multipe users
	for i := 0; i < users; i++ {
		var userName string
		var userAge int
		var ticketPrice int

		fmt.Printf("\n--- User %d ---\n", i)

		fmt.Print("Enter user name:")
		fmt.Scan(&userName)

		fmt.Print("Enter user age:")
		fmt.Scan(&userAge)

		//Validation
		if userAge <= 0 {
			fmt.Print("Invalid age entered ❌")
			return
		}

		fmt.Println("----- Ticket Details -----")

		//Decision making with if-else
		if userAge >= 18 {
			ticketPrice = 1000
			fmt.Printf("Hello %s , you are an Adult ✅\n", userName)
			fmt.Println("Ticket Type: Adult")
		} else if userAge >= 13 {
			ticketPrice = 700
			fmt.Printf("Hello %s , you are a Teen ⚠️\n", userName)
			fmt.Println("Ticket Type: Teen")
			fmt.Println("Parental guidance is recommended.")
		} else {
			ticketPrice = 500
			fmt.Printf("Hello %s , you are a Child 👶\n", userName)
			fmt.Println("Ticket Type: Child")
			fmt.Println("Parental guidance is required.")
		}
		// Add to total revenue
		totalRevenue += ticketPrice
	}
	fmt.Println("\n==========================")
	fmt.Println("Total Revenue:", totalRevenue)
	fmt.Println("==========================")
}
