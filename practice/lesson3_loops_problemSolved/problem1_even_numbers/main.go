package main

import "fmt"

func main() {
	count := 0
	fmt.Println("Even numbers.")
	//Loop for multiple users
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			count++
		}
	}

	fmt.Println("Total even numbers:", count)
}
