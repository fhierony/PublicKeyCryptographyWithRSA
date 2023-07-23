package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var num, pow, mod int64

		fmt.Print("Enter num : ")
		if _, err := fmt.Scanln(&num); err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Print("Enter pow : ")
		if _, err := fmt.Scanln(&pow); err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Print("Enter mod : ")
		if _, err := fmt.Scanln(&mod); err != nil {
			fmt.Println("Error: ", err)
		}

		if num < 0 || pow < 0 || mod < 0 {
			break
		}

		numPowMathPow := math.Pow(float64(num), float64(pow))
		fmt.Println("num ^ pow")
		fmt.Println("fastExp() : ", fastExp(num, pow))
		fmt.Println("math.Pow() : ", int64(numPowMathPow))

		fmt.Println("num ^ pow % mod")
		fmt.Println("fastExpMod() : ", fastExpMod(num, pow, mod))
		fmt.Println("math.Pow() : ", int64(numPowMathPow)%mod)
	}
}

// Use fast exponentiation to calculate num ^ pow.
func fastExp(num, pow int64) int64 {
	var result int64 = 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

func fastExpMod(num, pow, mod int64) int64 {
	var result int64 = 1
	for pow > 0 {
		if pow%2 == 1 {
			result = result * num % mod
		}
		pow /= 2
		num *= num % mod
	}
	return result
}
