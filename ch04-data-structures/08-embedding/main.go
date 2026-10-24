// Chapter 4.8: Embedding (composition instead of inheritance)
// Run: go run ./ch04-data-structures/08-embedding
package main

import (
	"fmt"
	"time"
)

// Shared pieces
type Timestamps struct {
	Created time.Time
	Updated time.Time
}

func (t *Timestamps) Touch(now time.Time) { t.Updated = now }

type Contact struct {
	Email, Phone string
}

func (c Contact) Reach() string { return "email " + c.Email + " or call " + c.Phone }

// Employee is built from parts, like a car from an engine and wheels.
type Employee struct {
	Name string
	Contact
	Timestamps
}

// Supplier reuses the same parts.
type Supplier struct {
	Company string
	Contact
}

func main() {
	day1 := time.Date(2026, 1, 5, 9, 0, 0, 0, time.UTC)
	e := Employee{
		Name:       "Gita",
		Contact:    Contact{"gita@shop.in", "98000 00001"},
		Timestamps: Timestamps{Created: day1, Updated: day1},
	}

	// Fields and methods of embedded types are "promoted".
	fmt.Println(e.Name, "-", e.Reach())
	fmt.Println("email directly:", e.Email)

	e.Touch(day1.AddDate(0, 2, 0))
	fmt.Println("updated:", e.Updated.Format("2 Jan 2006"))

	s := Supplier{"Fresh Farms", Contact{"orders@freshfarms.in", "080 1234 5678"}}
	fmt.Println(s.Company, "-", s.Reach())

	// Both types can be used wherever a "Reacher" is needed.
	type Reacher interface{ Reach() string }
	for _, r := range []Reacher{e, s} {
		fmt.Println("->", r.Reach())
	}
}
