// Chapter 2.5: the for loop
// Run: go run ./ch02-control-flow/05-for-loop
package main

import "fmt"

func main() {
	// init; condition; post
	for i := 1; i <= 5; i++ {
		fmt.Println("Lap", i)
	}

	// Counting down
	for n := 3; n > 0; n-- {
		fmt.Println(n, "...")
	}
	fmt.Println("Lift off!")

	// Step by more than one
	for year := 2020; year <= 2032; year += 4 {
		fmt.Println("Olympic year:", year)
	}

	// A running total: savings with 1% monthly interest
	savings := 0.0
	for month := 1; month <= 12; month++ {
		savings += 1000
		savings *= 1.01
	}
	fmt.Printf("After 12 months: Rs %.2f\n", savings)

	// Go 1.22+: range over a number counts 0..n-1
	for i := range 3 {
		fmt.Println("range over int:", i)
	}
}
