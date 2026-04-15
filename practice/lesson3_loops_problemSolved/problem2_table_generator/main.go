package main

import "fmt"

func main() {
	var num int

	// User input
	fmt.Print("Enter a number:")
	fmt.Scan(&num)

	fmt.Println("\nMultiplication Table of", num)
	fmt.Println("=========================")

	// Table Print(1 to 20)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
