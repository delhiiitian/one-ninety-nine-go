// Chapter 7.6: Unit tests
// Run the program: go run ./ch07-modules-tooling/06-unit-tests
// Run the tests:   go test -v ./ch07-modules-tooling/06-unit-tests
package main

import "fmt"

func main() {
	fmt.Println("Rs 999 with 20% off:", Discount(999, 20))
	fmt.Println("Now run: go test -v ./ch07-modules-tooling/06-unit-tests")
}
