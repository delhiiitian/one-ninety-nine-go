// Chapter 3.4: Variadic functions (any number of arguments)
// Run: go run ./ch03-functions/04-variadic
package main

import (
	"fmt"
	"strings"
)

// ...float64 means "zero or more float64 values"; inside, prices is a slice.
func total(prices ...float64) float64 {
	sum := 0.0
	for _, p := range prices {
		sum += p
	}
	return sum
}

// Fixed parameters come first, the variadic one must be last.
func announce(title string, names ...string) {
	if len(names) == 0 {
		fmt.Println(title + ": nobody yet")
		return
	}
	fmt.Println(title+":", strings.Join(names, ", "))
}

func main() {
	fmt.Println(total())
	fmt.Println(total(40))
	fmt.Println(total(40, 25.5, 99))

	// Already have a slice? Spread it with ...
	cart := []float64{120, 80, 45}
	fmt.Println("cart total:", total(cart...))

	announce("Winners")
	announce("Winners", "Kavya", "Rohan", "Sam")

	// fmt.Println itself is variadic!
	fmt.Println("one", 2, 3.0, true)
}
