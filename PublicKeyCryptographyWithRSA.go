package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const numTests = 20

func main() {
	testKnownValues()

	fmt.Println()

	for {
		var digits int

		fmt.Print("Enter digits : ")
		if _, err := fmt.Scanln(&digits); err != nil {
			fmt.Println("Error: ", err)
		}

		if digits < 1 {
			break
		}

		min := int64(math.Pow(float64(10), float64(digits-1)))
		max := min * 10

		if min == 1 {
			min = 2
		}
		fmt.Printf("Prime: %d\n\n", findPrime(min, max, numTests))
	}
}

// Return a pseudo random number in the range [min, max).
func randRange(min, max int64) int64 {
	return min + random.Int63n(max-min)
}

func isProbablyPrime(p int64, numTests int) bool {
	for i := 0; i < numTests; i++ {
		n := randRange(2, p)
		result := fastExpMod(n, p-1, p)
		if result != 1 {
			return false
		}
	}

	return true
}

func fastExpMod(num, pow, mod int64) int64 {
	var result int64 = 1
	for pow > 0 {
		if pow%2 == 1 {
			result = (result * num) % mod
		}
		pow /= 2
		num = (num * num) % mod
	}
	return result
}

func findPrime(min, max int64, numTests int) int64 {
	for {
		n := randRange(min, max+1)
		if n%2 == 0 {
			continue
		}

		if isProbablyPrime(n, numTests) {
			return n
		}
	}
}

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf("Probability: %f%%\n", (1-1/math.Pow(float64(2), float64(numTests)))*100)
	fmt.Println()

	fmt.Println("Primes:")
	for _, prime := range primes {
		fmt.Printf("%d", prime)
		if isProbablyPrime(int64(prime), numTests) {
			fmt.Printf("\tPrime")
		}
		fmt.Println()
	}

	fmt.Println()

	fmt.Println("Composites:")
	for _, composite := range composites {
		fmt.Printf("%d", composite)
		if !isProbablyPrime(int64(composite), numTests) {
			fmt.Printf("\tComposite")
		}
		fmt.Println()
	}
}
