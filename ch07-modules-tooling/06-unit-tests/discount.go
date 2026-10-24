package main

// Discount returns the price after a percentage discount.
// Percent outside 0..100 is clamped so the price never goes negative.
func Discount(price, percent float64) float64 {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	return price * (100 - percent) / 100
}
