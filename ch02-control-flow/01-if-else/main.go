// Chapter 2.1: if / else
// Run: go run ./ch02-control-flow/01-if-else
package main

import "fmt"

func main() {
	temperature := 34

	// A single decision
	if temperature > 30 {
		fmt.Println("It's hot. Carry water.")
	}

	// Two paths
	balance := 1500
	price := 2000
	if balance >= price {
		fmt.Println("Payment approved.")
	} else {
		fmt.Println("Payment declined: need", price-balance, "more.")
	}

	// Many paths: checked from top to bottom, first match wins.
	marks := 72
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: Needs improvement")
	}

	// Combine conditions with && and ||
	age, hasLicence := 19, true
	if age >= 18 && hasLicence {
		fmt.Println("You may drive.")
	}
}
