package main

import "fmt"

func main() {
	// Variable to hold the running sum
	total := 0
	// Iterate over values below 1000
	for i := 0; i < 1000; i++ {
		// If divisible by 3 or 5, add it to the running total
		if i % 3 == 0 || i % 5 == 0 {
			total += i
		}
	}

	fmt.Println(total)
}
