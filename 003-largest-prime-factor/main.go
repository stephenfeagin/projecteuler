// THIS CODE WILL NOT WORK FOR THE PROJECT EULER INPUT
// This is not sufficiently optimized to deal with as big a number
// as Project Euler asks for. I will probably need to re-do this
// using the `big` package

package main

import "fmt"

// sieve returns all prime numbers < n
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

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func largestPrimeFactor(n int) int {
	primes := sieve((n / 2) + 1)
	for i := len(primes) - 1; i >= 0; i-- {
		if n%primes[i] == 0 {
			return primes[i]
		}
	}
	return 0
}

func main() {
	// fmt.Println(largestPrimeFactor(600851475143))
	fmt.Println(sieve(600851475143))
}
