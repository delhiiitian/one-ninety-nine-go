// Chapter 2.6: for as a "while" loop and the endless loop
// Run: go run ./ch02-control-flow/06-for-as-while
package main

import "fmt"

func main() {
	// Only a condition: runs while it is true (other languages call this "while").
	water := 100 // litres in the tank
	day := 0
	for water > 0 {
		water -= 30
		day++
	}
	fmt.Println("Tank empty after", day, "days")

	// How many times can you halve 1000 before reaching 1?
	n, steps := 1000, 0
	for n > 1 {
		n /= 2
		steps++
	}
	fmt.Println("Halvings:", steps)

	// No condition at all: loops forever until break or return.
	// Real servers and games run like this.
	tick := 0
	for {
		tick++
		if tick == 3 {
			fmt.Println("Stopping at tick", tick)
			break
		}
		fmt.Println("tick", tick)
	}
}
