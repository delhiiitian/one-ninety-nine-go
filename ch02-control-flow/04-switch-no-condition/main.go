// Chapter 2.4: switch with no condition (a cleaner if-else chain)
// Run: go run ./ch02-control-flow/04-switch-no-condition
package main

import "fmt"

func main() {
	for _, units := range []int{80, 250, 520} {
		fmt.Printf("%d units -> Rs %.2f\n", units, electricityBill(units))
	}

	for _, hour := range []int{7, 13, 19, 23} {
		fmt.Printf("%02d:00 -> %s\n", hour, greeting(hour))
	}
}

// electricityBill uses slab rates: each case is a full condition.
func electricityBill(units int) float64 {
	switch {
	case units <= 100:
		return float64(units) * 3.0
	case units <= 300:
		return 100*3.0 + float64(units-100)*5.0
	default:
		return 100*3.0 + 200*5.0 + float64(units-300)*7.5
	}
}

func greeting(hour int) string {
	switch {
	case hour < 12:
		return "Good morning"
	case hour < 17:
		return "Good afternoon"
	case hour < 21:
		return "Good evening"
	default:
		return "Good night"
	}
}
