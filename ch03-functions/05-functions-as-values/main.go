// Chapter 3.5: Functions as values
// Run: go run ./ch03-functions/05-functions-as-values
package main

import (
	"fmt"
	"strings"
)

// A function type: any function that takes a float64 and returns a float64.
type Discount func(price float64) float64

func noDiscount(p float64) float64  { return p }
func festiveSale(p float64) float64 { return p * 0.8 } // 20% off
func flatHundred(p float64) float64 { return max(p-100, 0) }

// checkout accepts a function as a parameter.
func checkout(price float64, apply Discount) float64 {
	return apply(price)
}

// apply runs fn on every item and returns the new list.
func apply(words []string, fn func(string) string) []string {
	out := make([]string, 0, len(words))
	for _, w := range words {
		out = append(out, fn(w))
	}
	return out
}

func main() {
	offers := map[string]Discount{
		"none":    noDiscount,
		"festive": festiveSale,
		"flat100": flatHundred,
	}
	for _, name := range []string{"none", "festive", "flat100"} {
		fmt.Printf("%-8s -> Rs %.2f\n", name, checkout(750, offers[name]))
	}

	// An anonymous function, written right where it is used.
	double := func(p float64) float64 { return p * 2 }
	fmt.Println("double:", checkout(750, double))

	fmt.Println(apply([]string{"go", "is", "fun"}, strings.ToUpper))
	fmt.Println(apply([]string{"go", "is", "fun"}, func(s string) string { return s + "!" }))
}
