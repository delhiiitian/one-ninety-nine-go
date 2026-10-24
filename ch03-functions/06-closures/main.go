// Chapter 3.6: Closures (functions that remember)
// Run: go run ./ch03-functions/06-closures
package main

import "fmt"

// tokenMachine returns a function. Each call to that function remembers "next".
func tokenMachine(prefix string) func() string {
	next := 0
	return func() string {
		next++
		return fmt.Sprintf("%s-%03d", prefix, next)
	}
}

// A closure that remembers a setting: a price formatter for a currency.
func formatter(symbol string) func(float64) string {
	return func(amount float64) string {
		return fmt.Sprintf("%s%.2f", symbol, amount)
	}
}

func main() {
	pharmacy := tokenMachine("PH")
	bank := tokenMachine("BK")

	fmt.Println(pharmacy()) // PH-001
	fmt.Println(pharmacy()) // PH-002
	fmt.Println(bank())     // BK-001  (its own counter)
	fmt.Println(pharmacy()) // PH-003

	rupees := formatter("Rs ")
	dollars := formatter("$")
	fmt.Println(rupees(1499), dollars(17.5))

	// Closures can change variables from the surrounding function.
	visits := 0
	recordVisit := func() { visits++ }
	recordVisit()
	recordVisit()
	fmt.Println("visits:", visits)
}
