// Chapter 2.7: range
// Run: go run ./ch02-control-flow/07-range
package main

import "fmt"

func main() {
	// Over a slice: index and value
	cart := []string{"rice", "dal", "oil"}
	for i, item := range cart {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Ignore the index with _
	prices := []float64{65, 120, 180}
	total := 0.0
	for _, p := range prices {
		total += p
	}
	fmt.Println("total:", total)

	// Over a map: key and value (order is not guaranteed!)
	capitals := map[string]string{"India": "New Delhi", "Japan": "Tokyo"}
	for country, capital := range capitals {
		fmt.Println(country, "->", capital)
	}

	// Over a string: runes (characters), with byte positions
	for pos, ch := range "नमस्ते" {
		fmt.Printf("%d:%c ", pos, ch)
	}
	fmt.Println()

	// Only the index
	for i := range cart {
		cart[i] = "organic " + cart[i]
	}
	fmt.Println(cart)
}
