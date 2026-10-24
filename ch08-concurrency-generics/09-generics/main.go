// Chapter 8.9: Generics (one function for many types)
// Run: go run ./ch08-concurrency-generics/09-generics
package main

import (
	"cmp"
	"fmt"
	"strings"
)

// Number is a constraint: the set of types a type parameter may be.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum works for ints, float64s, and types built on them.
func Sum[T Number](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Map applies f to every element: from a slice of T to a slice of U.
func Map[T, U any](items []T, f func(T) U) []U {
	out := make([]U, 0, len(items))
	for _, it := range items {
		out = append(out, f(it))
	}
	return out
}

// Filter keeps the elements for which keep returns true.
func Filter[T any](items []T, keep func(T) bool) []T {
	var out []T
	for _, it := range items {
		if keep(it) {
			out = append(out, it)
		}
	}
	return out
}

// MaxBy finds the largest element by a key. cmp.Ordered allows < and >.
func MaxBy[T any, K cmp.Ordered](items []T, key func(T) K) T {
	best := items[0]
	for _, it := range items[1:] {
		if key(it) > key(best) {
			best = it
		}
	}
	return best
}

// Stack is a generic type: a stack of plates of any kind.
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

type Rupees float64 // ~float64 in Number lets this work with Sum

type Product struct {
	Name  string
	Price Rupees
}

func main() {
	fmt.Println("sum ints:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("sum floats:", Sum([]float64{2.5, 0.5}))

	cart := []Product{{"Pen", 20}, {"Bag", 899}, {"Book", 199}, {"Bottle", 349}}
	prices := Map(cart, func(p Product) Rupees { return p.Price })
	fmt.Println("cart total:", Sum(prices))

	names := Map(cart, func(p Product) string { return strings.ToUpper(p.Name) })
	fmt.Println("names:", names)

	cheap := Filter(cart, func(p Product) bool { return p.Price < 300 })
	fmt.Println("under 300:", cheap)

	fmt.Println("most expensive:", MaxBy(cart, func(p Product) Rupees { return p.Price }).Name)
	fmt.Println("longest name:", MaxBy(cart, func(p Product) int { return len(p.Name) }).Name)

	var undo Stack[string]
	undo.Push("type 'hello'")
	undo.Push("make bold")
	last, _ := undo.Pop()
	fmt.Println("undo:", last)

	var nums Stack[int]
	_, ok := nums.Pop()
	fmt.Println("pop from empty int stack ok?", ok)
}
