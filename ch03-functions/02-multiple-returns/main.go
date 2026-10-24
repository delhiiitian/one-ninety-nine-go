// Chapter 3.2: Multiple return values
// Run: go run ./ch03-functions/02-multiple-returns
package main

import (
	"errors"
	"fmt"
)

// Return two results at once: quotient and remainder.
func divide(total, groupSize int) (int, int) {
	return total / groupSize, total % groupSize
}

// The most common Go pattern: (result, error).
func withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("amount must be positive")
	}
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	return balance - amount, nil // nil means "no error"
}

func minMax(nums []int) (int, int) {
	lo, hi := nums[0], nums[0]
	for _, n := range nums {
		if n < lo {
			lo = n
		}
		if n > hi {
			hi = n
		}
	}
	return lo, hi
}

func main() {
	buses, left := divide(130, 40)
	fmt.Println("Full buses:", buses, "Students left over:", left)

	balance := 5000.0
	balance, err := withdraw(balance, 1200)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Balance:", balance)

	_, err = withdraw(balance, 99999)
	fmt.Println("Big withdrawal:", err)

	lo, hi := minMax([]int{31, 28, 35, 22, 40})
	fmt.Println("Coolest:", lo, "Hottest:", hi)
}
