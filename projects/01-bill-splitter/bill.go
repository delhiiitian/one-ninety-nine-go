package main

import "math"

// Settings that never change while the program runs.
const (
	TaxRate           = 0.05 // 5% GST on restaurant food
	DefaultTipPercent = 10.0
)

// Bill holds every number on the receipt.
type Bill struct {
	Subtotal   float64
	Tax        float64
	TipPercent float64
	Tip        float64
	Total      float64
	People     int
	PerPerson  float64
}

// Calculate does the maths only: no printing, no reading.
// Keeping it pure makes it easy to test and reuse.
func Calculate(subtotal, tipPercent float64, people int) Bill {
	tax := subtotal * TaxRate
	tip := subtotal * tipPercent / 100
	total := subtotal + tax + tip
	return Bill{
		Subtotal:   subtotal,
		Tax:        tax,
		TipPercent: tipPercent,
		Tip:        tip,
		Total:      total,
		People:     people,
		PerPerson:  roundUp(total/float64(people), 2),
	}
}

// roundUp rounds up to the given decimal places so the group never pays less than the total.
func roundUp(value float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Ceil(value*factor) / factor
}
