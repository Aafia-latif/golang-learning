package main

import "fmt"

func main() {
	var num int

	fmt.Print("How many many numbers do you want to enter?")
	fmt.Scan(&num)

	var sum int
	var num2 int

	// Loop for numbers input
	for i := 0; i < num; i++ {
		fmt.Print("Enter number:")
		fmt.Scan(&num2)
		sum += num2
	}

	// Calculate average
	average := float64(sum) / float64(num)

	// Output
	fmt.Println("\n-----Result-----")
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)
}
