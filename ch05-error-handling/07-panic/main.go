// Chapter 5.7: panic (for bugs, not for normal errors)
// Run: go run ./ch05-error-handling/07-panic
// This program crashes ON PURPOSE at the end. Read the output!
package main

import "fmt"

type Matrix struct {
	rows, cols int
	data       []float64
}

// NewMatrix panics if called with impossible sizes: that is a programmer mistake.
func NewMatrix(rows, cols int) *Matrix {
	if rows <= 0 || cols <= 0 {
		panic(fmt.Sprintf("NewMatrix: invalid size %dx%d", rows, cols))
	}
	return &Matrix{rows, cols, make([]float64, rows*cols)}
}

// MustPositive is a common naming style: "Must" functions panic instead of returning errors.
func MustPositive(n int) int {
	if n <= 0 {
		panic("MustPositive: got " + fmt.Sprint(n))
	}
	return n
}

func main() {
	m := NewMatrix(2, 3)
	fmt.Println("matrix created:", m.rows, "x", m.cols)

	fmt.Println("MustPositive(5) =", MustPositive(5))

	// The Go runtime panics on bugs like this one:
	defer fmt.Println("deferred calls still run during a panic")
	scores := []int{90, 85, 70}
	i := len(scores) // off-by-one bug: last index is len-1
	fmt.Println("about to read index", i)
	fmt.Println(scores[i]) // panic: index out of range
	fmt.Println("this line never runs")
}
