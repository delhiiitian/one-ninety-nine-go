// Chapter 4.1: Arrays (fixed size)
// Run: go run ./ch04-data-structures/01-arrays
package main

import "fmt"

func main() {
	// An egg tray holds exactly 6. The size is part of the type: [6]string.
	var tray [6]string
	tray[0] = "brown"
	tray[1] = "white"
	fmt.Println("tray:", tray, "length:", len(tray))

	// Array literal
	weekTemps := [7]float64{31, 33, 30, 29, 34, 35, 32}
	sum := 0.0
	for _, t := range weekTemps {
		sum += t
	}
	fmt.Printf("weekly average: %.1f\n", sum/float64(len(weekTemps)))

	// Let the compiler count with [...]
	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes, "count:", len(primes))

	// Arrays are values: assigning copies every element.
	backup := primes
	backup[0] = 999
	fmt.Println("original:", primes[0], "backup:", backup[0])

	// A 2D array: a tic-tac-toe board
	board := [3][3]rune{
		{'X', 'O', 'X'},
		{' ', 'X', 'O'},
		{'O', ' ', 'X'},
	}
	for _, row := range board {
		fmt.Printf("|%c|%c|%c|\n", row[0], row[1], row[2])
	}

	// tray[6] = "blue" // compile error: index out of range
}
