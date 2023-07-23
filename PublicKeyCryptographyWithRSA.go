package main

import (
	"fmt"
	"time"
)

func main() {
	var max int
	fmt.Printf("Max: ")
	if _, err := fmt.Scan(&max); err != nil {
		fmt.Println("Error: ", err)
	}

	start := time.Now()
	sieve := sieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		printSieve(sieve)

		primes := sieveToPrimes(sieve)
		fmt.Println(primes)
	}
}

// Build a sieve of Eratosthenes.
func sieveOfEratosthenes(max int) []bool {
	prime := make([]bool, max+1)

	if max < 2 {
		return prime
	}

	prime[2] = true
	for i := 3; i <= max; i += 2 {
		prime[i] = true
	}

	for i := 3; i*i <= max; i++ {
		if prime[i] {
			for j := i * 3; j <= max; j += i {
				prime[j] = false
			}
		}
	}

	return prime
}

func printSieve(sieve []bool) {
	if len(sieve) > 2 {
		fmt.Printf("2 ")
	}
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func sieveToPrimes(sieve []bool) interface{} {
	var primes []int

	if len(sieve) > 2 {
		primes = append(primes, 2)
	}
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
