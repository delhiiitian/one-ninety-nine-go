// Chapter 5.2: Creating your own errors
// Run: go run ./ch05-error-handling/02-creating-errors
package main

import (
	"errors"
	"fmt"
)

// errors.New for a fixed message
func checkAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age < 18 {
		return fmt.Errorf("must be 18 or older, got %d", age) // fmt.Errorf adds details
	}
	return nil
}

func bookSeats(available, wanted int) (int, error) {
	if wanted <= 0 {
		return available, errors.New("ask for at least one seat")
	}
	if wanted > available {
		return available, fmt.Errorf("only %d seats left, you asked for %d", available, wanted)
	}
	return available - wanted, nil
}

func main() {
	for _, age := range []int{25, 16, -3} {
		if err := checkAge(age); err != nil {
			fmt.Println("rejected:", err)
		} else {
			fmt.Println("age", age, "accepted")
		}
	}

	seats := 5
	for _, want := range []int{2, 4, 0, 3} {
		left, err := bookSeats(seats, want)
		if err != nil {
			fmt.Println("booking failed:", err)
			continue
		}
		seats = left
		fmt.Printf("booked %d, %d left\n", want, seats)
	}
}
