// Chapter 6.8: JSON (the language of web APIs)
// Run: go run ./ch06-files-io/08-json
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Struct tags tell encoding/json what the keys are called.
type Order struct {
	ID       int      `json:"id"`
	Customer string   `json:"customer"`
	Items    []Item   `json:"items"`
	Paid     bool     `json:"paid"`
	Coupon   string   `json:"coupon,omitempty"` // left out when empty
	Notes    []string `json:"-"`                // never written
}

type Item struct {
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}

func main() {
	order := Order{
		ID: 101, Customer: "Sana",
		Items: []Item{{"Pizza", 2, 349}, {"Cola", 2, 60}},
		Paid:  true,
		Notes: []string{"internal only"},
	}

	// Go value -> JSON text (Marshal)
	data, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	path := filepath.Join(os.TempDir(), "order.json")
	os.WriteFile(path, data, 0o644)
	defer os.Remove(path)

	// JSON text -> Go value (Unmarshal)
	raw, _ := os.ReadFile(path)
	var loaded Order
	if err := json.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("bad JSON:", err)
		return
	}
	total := 0.0
	for _, it := range loaded.Items {
		total += float64(it.Qty) * it.Price
	}
	fmt.Printf("loaded order %d for %s: total Rs %.2f\n", loaded.ID, loaded.Customer, total)

	// Unknown shapes: decode into a map
	var anything map[string]any
	json.Unmarshal([]byte(`{"city":"Pune","temp":31.5,"rain":false}`), &anything)
	fmt.Println("city:", anything["city"], "temp:", anything["temp"])

	// Errors tell you where the JSON is broken
	err = json.Unmarshal([]byte(`{"id": "not-a-number"}`), &loaded)
	fmt.Println("error:", err)
}
