// Chapter 1.4: Constants and iota
// Run: go run ./ch01-basics/04-constants-iota
package main

import "fmt"

// A constant never changes while the program runs.
const Pi = 3.14159
const AppName = "ShopEasy"

// Group related constants together.
const (
	MaxLoginTries = 3
	TaxRate       = 0.18 // 18% GST
)

// iota counts up from 0 inside a const block: perfect for "kinds" of things.
type OrderStatus int

const (
	Pending   OrderStatus = iota // 0
	Paid                         // 1
	Shipped                      // 2
	Delivered                    // 3
)

// iota can drive a formula, too: file sizes.
const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20
	GB                    // 1 << 30
)

func main() {
	radius := 2.0
	fmt.Println(AppName, "circle area:", Pi*radius*radius)

	price := 500.0
	fmt.Println("price with tax:", price*(1+TaxRate))
	fmt.Println("max login tries:", MaxLoginTries)

	status := Shipped
	fmt.Println("order status code:", status)
	fmt.Println("is delivered?", status == Delivered)

	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB)

	// Pi = 3 // would not compile: cannot assign to a constant
}
