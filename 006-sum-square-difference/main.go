// Go doesn't have things like vectorization or generator expressions, so
// we do things more manually. More verbose, less magic, but very clear.
package main

import "fmt"

// For every number in the slice, add its square to a running total
func sumOfSquares(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num * num
	}
	return total
}

// Add every number in the slice to a running total, then
// return the square of that total.
func squareOfSums(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total * total
}

func main() {
	// Start by making the sequence of 1 to 100
	seq := make([]int, 100)
	for i := 0; i < 100; i++ {
		seq[i] = i + 1
	}
	// Then take the difference between the two values
	fmt.Println(squareOfSums(seq) - sumOfSquares(seq))
}
