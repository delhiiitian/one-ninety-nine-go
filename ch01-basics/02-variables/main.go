// Chapter 1.2: Variables
// Run: go run ./ch01-basics/02-variables
package main

import "fmt"

func main() {
	// 1. Full form: var name type = value
	var city string = "Bengaluru"

	// 2. Let Go infer the type from the value.
	var population = 13_600_000

	// 3. Short form (only inside functions): name := value
	temperature := 27.5

	// 4. Declare now, assign later. Go gives a "zero value" first.
	var visitors int // starts at 0
	fmt.Println("visitors before:", visitors)
	visitors = 1200

	// 5. Several at once.
	var x, y = 10, 20
	name, age := "Asha", 31

	// 6. Change a value (the type must stay the same).
	temperature = temperature + 1.5

	fmt.Println("city:", city)
	fmt.Println("population:", population)
	fmt.Println("temperature:", temperature)
	fmt.Println("visitors after:", visitors)
	fmt.Println("x + y =", x+y)
	fmt.Println(name, "is", age, "years old")

	// Swapping is a one-liner in Go.
	x, y = y, x
	fmt.Println("after swap x =", x, "y =", y)
}
