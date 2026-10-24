// Chapter 1.5: Operators
// Run: go run ./ch01-basics/05-operators
package main

import "fmt"

func main() {
	// Arithmetic
	a, b := 17, 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b, "(whole-number division drops the decimal)")
	fmt.Println("a % b =", a%b, "(remainder)")
	fmt.Println("17.0 / 5.0 =", 17.0/5.0)

	// Assignment shortcuts
	score := 10
	score += 5 // score = score + 5
	score *= 2 // score = score * 2
	score--    // score = score - 1 (a statement, not an expression)
	fmt.Println("score:", score)

	// Comparison: the result is a bool
	age := 20
	fmt.Println("age >= 18:", age >= 18)
	fmt.Println("age == 21:", age == 21)
	fmt.Println("age != 21:", age != 21)

	// Logic: && (and), || (or), ! (not)
	hasTicket := true
	isVIP := false
	fmt.Println("can enter:", age >= 18 && hasTicket)
	fmt.Println("fast lane:", isVIP || age > 60)
	fmt.Println("needs ticket:", !hasTicket)

	// Precedence: * and / before + and -. Use brackets to be clear.
	fmt.Println("2 + 3 * 4 =", 2+3*4)
	fmt.Println("(2 + 3) * 4 =", (2+3)*4)
}
