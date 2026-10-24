// Chapter 3.3: Named return values
// Run: go run ./ch03-functions/03-named-returns
package main

import "fmt"

// Named results document what comes back and start at their zero values.
func splitSalary(gross float64) (tax, pf, takeHome float64) {
	tax = gross * 0.10
	pf = gross * 0.12
	takeHome = gross - tax - pf
	return // a "naked" return sends back the named values
}

// Explicit returns still work with named results, and are clearer in long functions.
func bmi(weightKg, heightM float64) (value float64, category string) {
	value = weightKg / (heightM * heightM)
	switch {
	case value < 18.5:
		return value, "Underweight"
	case value < 25:
		return value, "Normal"
	default:
		return value, "Overweight"
	}
}

func main() {
	tax, pf, home := splitSalary(80000)
	fmt.Printf("Tax: %.0f  PF: %.0f  Take-home: %.0f\n", tax, pf, home)

	v, cat := bmi(68, 1.72)
	fmt.Printf("BMI %.1f (%s)\n", v, cat)
}
