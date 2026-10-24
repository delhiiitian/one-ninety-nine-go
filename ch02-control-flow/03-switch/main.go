// Chapter 2.3: switch
// Run: go run ./ch02-control-flow/03-switch
package main

import "fmt"

func main() {
	day := "Sat"

	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri": // several values in one case
		fmt.Println(day, "is a working day. Shop opens at 9.")
	case "Sat":
		fmt.Println(day, "shop opens at 10.")
	case "Sun":
		fmt.Println(day, "shop is closed.")
	default:
		fmt.Println("Unknown day:", day)
	}

	// No "break" needed: Go stops after the matching case.
	// Use fallthrough only when you really want the next case to run too.
	plan := "gold"
	fmt.Print("Benefits for ", plan, ": ")
	switch plan {
	case "gold":
		fmt.Print("lounge access, ")
		fallthrough
	case "silver":
		fmt.Print("extra baggage, ")
		fallthrough
	case "basic":
		fmt.Println("seat selection")
	}

	// Switch can start with a short statement, like if.
	switch code := 404; code {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Page not found")
	case 500:
		fmt.Println("Server error")
	}
}
