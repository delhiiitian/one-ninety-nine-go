// Chapter 5.6: Sentinel errors (named, shared error values)
// Run: go run ./ch05-error-handling/06-sentinel-errors
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Package-level errors that callers can compare against.
var (
	ErrEmptyCart   = errors.New("cart is empty")
	ErrInvalidCode = errors.New("invalid coupon code")
	ErrExpiredCode = errors.New("coupon has expired")
)

func applyCoupon(cartTotal float64, code string) (float64, error) {
	if cartTotal == 0 {
		return 0, ErrEmptyCart
	}
	switch code {
	case "SAVE10":
		return cartTotal * 0.9, nil
	case "DIWALI25":
		return cartTotal, ErrExpiredCode
	default:
		return cartTotal, fmt.Errorf("%q: %w", code, ErrInvalidCode)
	}
}

func main() {
	tests := []struct {
		total float64
		code  string
	}{{1000, "SAVE10"}, {1000, "DIWALI25"}, {1000, "FREE"}, {0, "SAVE10"}}

	for _, t := range tests {
		total, err := applyCoupon(t.total, t.code)
		switch {
		case err == nil:
			fmt.Printf("%-8s new total Rs %.0f\n", t.code, total)
		case errors.Is(err, ErrExpiredCode):
			fmt.Printf("%-8s sorry, that offer ended\n", t.code)
		case errors.Is(err, ErrInvalidCode):
			fmt.Printf("%-8s check the code and try again (%v)\n", t.code, err)
		case errors.Is(err, ErrEmptyCart):
			fmt.Printf("%-8s add something to your cart first\n", t.code)
		}
	}

	// The most famous sentinel: io.EOF means "no more data", not a failure.
	r := bufio.NewReader(strings.NewReader("line one\nline two\n"))
	for {
		line, err := r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			fmt.Println("(end of input)")
			break
		}
		fmt.Print("read: ", line)
	}
}
