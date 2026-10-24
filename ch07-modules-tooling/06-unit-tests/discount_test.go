package main

import "testing"

// Test functions start with Test and take *testing.T. Run: go test ./ch07-modules-tooling/06-unit-tests
func TestDiscountHalfPrice(t *testing.T) {
	got := Discount(200, 50)
	if got != 100 {
		t.Errorf("Discount(200, 50) = %v, want 100", got)
	}
}

func TestDiscountNeverNegative(t *testing.T) {
	if got := Discount(200, 150); got != 0 {
		t.Errorf("Discount(200, 150) = %v, want 0", got)
	}
}

func TestDiscountIgnoresNegativePercent(t *testing.T) {
	if got := Discount(200, -10); got != 200 {
		t.Errorf("Discount(200, -10) = %v, want 200", got)
	}
}
