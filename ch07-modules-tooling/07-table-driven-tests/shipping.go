package main

import "errors"

var ErrTooHeavy = errors.New("parcel too heavy (max 30 kg)")

// ShippingCost charges by weight band, with free shipping on big orders.
func ShippingCost(weightKg, orderValue float64) (float64, error) {
	switch {
	case weightKg <= 0:
		return 0, errors.New("weight must be positive")
	case weightKg > 30:
		return 0, ErrTooHeavy
	case orderValue >= 999:
		return 0, nil
	case weightKg <= 1:
		return 40, nil
	case weightKg <= 5:
		return 80, nil
	default:
		return 80 + (weightKg-5)*15, nil
	}
}
