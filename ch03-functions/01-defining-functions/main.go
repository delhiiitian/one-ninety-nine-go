// Chapter 3.1: Defining and calling functions
// Run: go run ./ch03-functions/01-defining-functions
package main

import "fmt"

// No inputs, no output
func printLine() {
	fmt.Println("------------------------")
}

// Inputs (parameters) with their types
func greet(name string, hour int) {
	if hour < 12 {
		fmt.Println("Good morning,", name)
	} else {
		fmt.Println("Hello,", name)
	}
}

// Parameters of the same type can share it; the return type comes last.
func area(width, height float64) float64 {
	return width * height
}

// Functions can call other functions.
func paintCost(width, height, pricePerSqM float64) float64 {
	return area(width, height) * pricePerSqM
}

func main() {
	printLine()
	greet("Priya", 9)
	greet("Arjun", 15)
	printLine()

	room := area(4, 3.5)
	fmt.Println("Room area:", room, "sq m")
	fmt.Printf("Painting cost: Rs %.2f\n", paintCost(4, 3.5, 120))
	printLine()
}
