// Chapter 1.7: Type conversion
// Run: go run ./ch01-basics/07-type-conversion
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Go never converts types for you. You ask: Type(value).
	apples := 7
	people := 2
	share := float64(apples) / float64(people)
	fmt.Println("each person gets", share, "apples")

	// float -> int cuts off the decimal (no rounding).
	price := 99.99
	fmt.Println("int(price):", int(price))

	// Text <-> numbers use the strconv package.
	ageText := "42"
	age, err := strconv.Atoi(ageText) // "ASCII to integer"
	if err != nil {
		fmt.Println("not a number:", err)
		return
	}
	fmt.Println("next year:", age+1)

	fmt.Println("Itoa:", strconv.Itoa(2026)+" is a year")

	weight, _ := strconv.ParseFloat("72.5", 64)
	fmt.Println("weight + 0.5 =", weight+0.5)

	ok, _ := strconv.ParseBool("true")
	fmt.Println("parsed bool:", ok)

	// Bad input gives an error instead of crashing.
	_, err = strconv.Atoi("forty-two")
	fmt.Println("error example:", err)

	// Careful: string(65) gives "A" (a character), not "65".
	fmt.Println("string(rune(65)):", string(rune(65)))
	fmt.Println("strconv.Itoa(65):", strconv.Itoa(65))
}
