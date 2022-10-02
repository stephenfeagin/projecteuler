// Same approach as the R solution
package main

import "fmt"

func main() {
	// Keep track of the running total, starting with two because we start looping
	// after the first two terms
	total := 2
	// Keep track of the previous two values
	lastTwo := []int{1, 2}
	// Get the next term in the sequence by adding the previous two
	nextTerm := lastTwo[0] + lastTwo[1]
	for nextTerm < 4000000 {
		// If the new term is even, add it to the running total
		if nextTerm%2 == 0 {
			total += nextTerm
		}

		// Update the lastTwo slice
		lastTwo = []int{lastTwo[1], nextTerm}

		// Update the nextTerm value
		nextTerm = lastTwo[0] + lastTwo[1]
	}
	fmt.Println(total)
}
