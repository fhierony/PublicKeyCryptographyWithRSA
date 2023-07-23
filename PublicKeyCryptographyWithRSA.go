package main

import (
	"fmt"
	"time"
)

var primes []int

func main() {
	primes = sieveToPrimes(eulersSieve(20_000_000))

	for {
		var num int

		fmt.Print("Enter num : ")
		if _, err := fmt.Scanln(&num); err != nil {
			fmt.Println("Error: ", err)
		}

		if num < 2 {
			break
		}

		// Find the factors the slow way.
		start := time.Now()
		factors := findFactors(num)
		elapsed := time.Since(start)
		fmt.Printf("findFactors: %f seconds\n", elapsed.Seconds())
		fmt.Printf("Prime factors of %d are : %v\n", num, factors)
		fmt.Println("Multiplication of prime factors =", multiplySlice(factors))
		fmt.Println()

		// Use the Euler's sieve to find the factors.
		start = time.Now()
		factors = findFactorsSieve(num)
		elapsed = time.Since(start)
		fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
		fmt.Printf("Prime factors of %d are : %v\n", num, factors)
		fmt.Println("Multiplication of prime factors =", multiplySlice(factors))
		fmt.Println()
	}
}

func findFactors(num int) []int {
	var primeFactors []int

	for num%2 == 0 {
		primeFactors = append(primeFactors, 2)
		num /= 2
	}

	for i := 3; i*i <= num; i = i + 2 {
		for num%i == 0 {
			primeFactors = append(primeFactors, i)
			num /= i
		}
	}

	if num > 2 {
		primeFactors = append(primeFactors, num)
	}

	return primeFactors
}

func multiplySlice(primeFactors []int) int {
	num := 1
	for _, factor := range primeFactors {
		num *= factor
	}
	return num
}

func sieveToPrimes(sieve []bool) []int {
	var primesSlice []int

	if len(sieve) > 2 {
		primesSlice = append(primesSlice, 2)
	}
	for i := 3; i < len(sieve); i += 2 {
		if sieve[i] {
			primesSlice = append(primesSlice, i)
		}
	}

	return primesSlice
}

// Build an Euler’s Sieve.
func eulersSieve(max int) []bool {
	prime := make([]bool, max+1)

	if max < 2 {
		return prime
	}

	prime[2] = true
	for i := 3; i <= max; i += 2 {
		prime[i] = true
	}

	for i := 3; i*i <= max; i += 2 {
		if prime[i] {
			maxQuotient := max / i
			if maxQuotient%2 == 0 {
				maxQuotient--
			}
			for j := maxQuotient; j >= i; j -= 2 {
				if prime[j] {
					prime[i*j] = false
				}
			}
		}
	}

	return prime
}

func findFactorsSieve(num int) []int {
	var primeFactors []int

	for num%2 == 0 {
		primeFactors = append(primeFactors, 2)
		num /= 2
	}

	for i := 1; i < len(primes); i++ {
		prime := primes[i]
		for num%prime == 0 {
			primeFactors = append(primeFactors, prime)
			num /= prime
		}
	}

	if num > 2 {
		primeFactors = append(primeFactors, num)
	}

	return primeFactors
}
