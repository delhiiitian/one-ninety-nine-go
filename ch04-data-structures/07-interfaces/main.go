// Chapter 4.7: Interfaces (what a thing can do)
// Run: go run ./ch04-data-structures/07-interfaces
package main

import (
	"fmt"
	"math"
)

// Any type with these methods is a Shape. No "implements" keyword needed.
type Shape interface {
	Area() float64
	Name() string
}

type Circle struct{ R float64 }
type Square struct{ Side float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
func (c Circle) Name() string  { return "circle" }
func (s Square) Area() float64 { return s.Side * s.Side }
func (s Square) Name() string  { return "square" }

// A payment example: the checkout does not care how you pay.
type PaymentMethod interface {
	Pay(amount float64) string
}

type UPI struct{ ID string }
type Card struct{ Last4 string }
type Cash struct{}

func (u UPI) Pay(a float64) string  { return fmt.Sprintf("Paid Rs %.0f via UPI %s", a, u.ID) }
func (c Card) Pay(a float64) string { return fmt.Sprintf("Paid Rs %.0f with card ****%s", a, c.Last4) }
func (Cash) Pay(a float64) string   { return fmt.Sprintf("Paid Rs %.0f in cash", a) }

func checkout(total float64, m PaymentMethod) {
	fmt.Println(m.Pay(total))
}

func main() {
	shapes := []Shape{Circle{1}, Square{3}, Circle{2.5}}
	total := 0.0
	for _, s := range shapes {
		fmt.Printf("%-6s %.2f\n", s.Name(), s.Area())
		total += s.Area()
	}
	fmt.Printf("total area: %.2f\n", total)

	checkout(499, UPI{"meera@okbank"})
	checkout(1299, Card{"4242"})
	checkout(60, Cash{})

	// A type switch looks at the concrete type behind an interface.
	var m PaymentMethod = Card{"1111"}
	switch v := m.(type) {
	case Card:
		fmt.Println("card ending", v.Last4)
	case UPI:
		fmt.Println("UPI id", v.ID)
	}

	// any (the empty interface) holds a value of any type.
	things := []any{42, "hello", 3.5, true}
	for _, t := range things {
		fmt.Printf("%v is %T\n", t, t)
	}
}
