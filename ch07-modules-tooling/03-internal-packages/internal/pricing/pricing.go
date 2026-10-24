// Package pricing holds the shop's pricing rules.
// Because it lives under a folder named "internal", only code inside
// 03-internal-packages/ may import it. Outsiders cannot depend on it.
package pricing

// Tax rates per category (GST)
var rates = map[string]float64{
	"food":        0.05,
	"electronics": 0.18,
	"books":       0.0,
}

// WithTax returns the price including tax for the category.
func WithTax(price float64, category string) float64 {
	return price * (1 + rates[category])
}
