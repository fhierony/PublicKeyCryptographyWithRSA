package main

import "fmt"

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

func main() {
	for {
		var a, b int64

		fmt.Print("Enter A : ")
		if _, err := fmt.Scanln(&a); err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Print("Enter B : ")
		if _, err := fmt.Scanln(&b); err != nil {
			fmt.Println("Error: ", err)
		}

		if a < 0 || b < 0 {
			break
		}

		fmt.Printf("A = %d, B = %d, GCD(A,B) = %d, LCM(A,B) = %d\n", a, b, gcd(a, b), lcm(a, b))
	}
}
