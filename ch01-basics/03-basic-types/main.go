// Chapter 1.3: Basic types
// Run: go run ./ch01-basics/03-basic-types
package main

import (
	"fmt"
	"math"
)

func main() {
	// Whole numbers
	var items int = 42             // size depends on the machine (64-bit today)
	var smallTemp int8 = -40       // -128 to 127
	var bytesOnDisk uint64 = 5_000 // unsigned: never negative

	// Decimal numbers
	var price float64 = 199.99
	var ratio float32 = 0.75

	// True or false
	var isOpen bool = true

	// Text
	var shop string = "Corner Store"

	// A single character is a rune (a Unicode code point).
	var grade rune = 'A'
	// A byte is one raw 8-bit value.
	var b byte = 'z'

	fmt.Printf("%-12s %-8T %v\n", "items", items, items)
	fmt.Printf("%-12s %-8T %v\n", "smallTemp", smallTemp, smallTemp)
	fmt.Printf("%-12s %-8T %v\n", "bytesOnDisk", bytesOnDisk, bytesOnDisk)
	fmt.Printf("%-12s %-8T %v\n", "price", price, price)
	fmt.Printf("%-12s %-8T %v\n", "ratio", ratio, ratio)
	fmt.Printf("%-12s %-8T %v\n", "isOpen", isOpen, isOpen)
	fmt.Printf("%-12s %-8T %v\n", "shop", shop, shop)
	fmt.Printf("%-12s %-8T %v (%c)\n", "grade", grade, grade, grade)
	fmt.Printf("%-12s %-8T %v (%c)\n", "b", b, b, b)

	// Zero values: what a variable holds before you set it.
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Printf("zero values: %v %v %v %q\n", zi, zf, zb, zs)

	// Limits help you pick the right type.
	fmt.Println("largest int8:", math.MaxInt8)
	fmt.Println("largest int64:", math.MaxInt64)
}
