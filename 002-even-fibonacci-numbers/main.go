// Same approach as the R solution
package main

import "fmt"

func main() {
	// Keep track of the running total, starting with two because we start looping
	// after the first two terms
	total := 2
	// Keep track of the previous two values
	lastTwo := []int{1, 2}

	for {
		// The next term in the Fibonacci sequence is the sum of the previous two
		nextTerm := lastTwo[0] + lastTwo[1]

		// If we cross 4mil, don't count it and break out of the loop
		if nextTerm >= 4000000 {
			break
		}
		// Re-assign the lastTwo slice
		lastTwo = []int{lastTwo[1], nextTerm}

		// If the new term is even, add it to the running total
		if nextTerm%2 == 0 {
			total += nextTerm
		}
	}
	fmt.Println(total)
}
