// Chapter 5.3: Custom error types (errors that carry data)
// Run: go run ./ch05-error-handling/03-custom-error-types
package main

import "fmt"

// FieldError says which form field is wrong and why.
type FieldError struct {
	Field  string
	Reason string
}

// Any type with an Error() string method is an error.
func (e *FieldError) Error() string {
	return e.Field + ": " + e.Reason
}

// OutOfStockError carries numbers the caller can use.
type OutOfStockError struct {
	Item      string
	Requested int
	InStock   int
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("%s: requested %d, only %d in stock", e.Item, e.Requested, e.InStock)
}

func validateSignup(email, password string) error {
	if email == "" {
		return &FieldError{"email", "is required"}
	}
	if len(password) < 8 {
		return &FieldError{"password", "must be at least 8 characters"}
	}
	return nil
}

func order(item string, qty int, stock map[string]int) error {
	if have := stock[item]; qty > have {
		return OutOfStockError{item, qty, have}
	}
	return nil
}

func main() {
	err := validateSignup("a@b.in", "short")
	if fe, ok := err.(*FieldError); ok {
		fmt.Println("highlight field:", fe.Field, "| message:", fe.Reason)
	}

	stock := map[string]int{"pens": 3}
	err = order("pens", 5, stock)
	fmt.Println("error:", err)
	if oos, ok := err.(OutOfStockError); ok {
		fmt.Printf("offer the customer the %d we have\n", oos.InStock)
	}
}
