package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize a pseudorandom number generator.
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const numTests = 20

func main() {
	random.Seed(time.Now().UnixNano())

	p := findPrime(10_000, 50_000, numTests)
	q := findPrime(10_000, 50_000, numTests)
	n := p * q

	t := totient(p, q)
	e := randomExponent(t)
	d := inverseMod(e, t)

	fmt.Println("*** Public ***")
	fmt.Println("Public key modulus:", n)
	fmt.Println("Public key exponent e:", e)

	fmt.Println()
	fmt.Println("*** Private ***")
	fmt.Printf("Primes:\t%d, %d\n", p, q)
	fmt.Println("λ(n):", t)
	fmt.Println("d:", d)

	fmt.Println()

	for {
		var message int64

		fmt.Print("Message: ")
		if _, err := fmt.Scanln(&message); err != nil {
			fmt.Println("Error: ", err)
		}

		if message < 1 || message > n-1 {
			break
		}

		// Encryption
		cipherText := fastExpMod(message, e, n)
		fmt.Printf("Ciphertext: %d\n", cipherText)

		// Decryption
		plainText := fastExpMod(cipherText, d, n)
		fmt.Printf("Plaintext: %d\n\n", plainText)
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

// Calculate the totient function λ(n)
// where n = p * q and p and q are prime.
func totient(p, q int64) int64 {
	return lcm(p-1, q-1)
}

func gcd(a, b int64) int64 {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func lcm(a, b int64) int64 {
	return b / gcd(a, b) * a
}

// Pick a random exponent e in the range (2, λ_n)
// such that gcd(e, λ_n) = 1.
func randomExponent(lambdaN int64) int64 {
	for {
		e := randRange(2, lambdaN)
		if gcd(e, lambdaN) == 1 {
			return e
		}
	}
}

func inverseMod(a, mod int64) int64 {
	var t int64 = 0
	var newT int64 = 1
	r := mod
	newR := a

	for newR != 0 {
		quotient := r / newR
		t, newT = newT, t-quotient*newT
		r, newR = newR, r-quotient*newR
	}

	if r > 1 {
		panic("a is not invertible")
	}

	if t < 0 {
		t += mod
	}

	return t
}
