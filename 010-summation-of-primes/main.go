package main

import (
	"fmt"
)

/*
p = 2
for i = 4; i < n; i += 2 {
	markers[4] = 1
}
*/

func sieve(n int) []int {
	markers := make([]bool, n)
	// 0 and 1 are non-prime by definition
	markers[0] = true
	markers[1] = true

	// Take p as the "current" prime -- the largest prime we have yet confirmed
	// Iterate through values of p until p^2 > n
	// This is because any values > sqrt(n) will already be marked off or will
	// be prime
	for p := 2; p*p < n; p++ {
		if !markers[p] { // if the value is "unmarked", still a possible prime,
			// mark all of the multiples of p starting at 2p
			for i := p * 2; i < n; i += p {
				markers[i] = true // mark it as composite
			}
		}
	}

	// fill out a slice with the actual prime numbers
	var primes []int
	for p := 2; p < n; p++ {
		if !markers[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	primes := sieve(2000000)
	total := 0
	for _, p := range primes {
		total += p
	}
	fmt.Println(total)
}
