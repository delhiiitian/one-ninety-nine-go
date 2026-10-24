package main

import (
	"errors"
	"testing"
)

// One table of cases, one loop: adding a case is adding one line.
func TestShippingCost(t *testing.T) {
	tests := []struct {
		name    string
		weight  float64
		value   float64
		want    float64
		wantErr error
	}{
		{"light parcel", 0.5, 300, 40, nil},
		{"exactly 1 kg", 1, 300, 40, nil},
		{"medium parcel", 3, 300, 80, nil},
		{"heavy parcel", 10, 300, 155, nil},
		{"free over 999", 10, 1200, 0, nil},
		{"too heavy", 31, 5000, 0, ErrTooHeavy},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) { // a named sub-test per row
			got, err := ShippingCost(tc.weight, tc.value)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v, want %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("ShippingCost(%v, %v) = %v, want %v", tc.weight, tc.value, got, tc.want)
			}
		})
	}
}

func TestShippingRejectsZeroWeight(t *testing.T) {
	if _, err := ShippingCost(0, 100); err == nil {
		t.Error("want an error for zero weight")
	}
}
