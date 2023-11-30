package maths

import (
	"math"
)

// GeneratePrimes returns the n-th prime number.
// The param primes []int is intended to be a slice of primes already generated.
func GeneratePrimes(primes []int, n int) int {
	if len(primes) < 2 {
		primes = []int{2, 3}
	}
	if len(primes) >= n {
		return primes[n-1]
	}
	for i := primes[len(primes)-1] + 2; len(primes) < n; i += 2 {
		// Check if i is prime by checking if it divisible by any of the previous primes.
		// Stop at the square root of i
		for _, v := range primes {
			// not a prime, stop loop
			if i%v == 0 {
				break
			}
			if math.Sqrt(float64(i)) < float64(v) {
				primes = append(primes, i)
				break
			}
		}
	}
	return primes[n-1]
}
