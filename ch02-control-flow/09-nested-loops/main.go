// Chapter 2.9: nested loops and putting it all together
// Run: go run ./ch02-control-flow/09-nested-loops
package main

import "fmt"

func main() {
	// A times table: a loop inside a loop.
	fmt.Println("Times table 1-5")
	for row := 1; row <= 5; row++ {
		for col := 1; col <= 5; col++ {
			fmt.Printf("%4d", row*col)
		}
		fmt.Println()
	}

	// A seating chart: rows A-D, seats 1-6, aisle after seat 3.
	fmt.Println("\nCinema seats")
	for r := 'A'; r <= 'D'; r++ {
		for s := 1; s <= 6; s++ {
			fmt.Printf("%c%d ", r, s)
			if s == 3 {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

	// FizzBuzz: the classic interview warm-up uses everything in this chapter.
	fmt.Println("\nFizzBuzz 1-15")
	for i := 1; i <= 15; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
