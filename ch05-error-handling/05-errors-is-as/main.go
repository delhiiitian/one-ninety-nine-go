// Chapter 5.5: errors.Is and errors.As
// Run: go run ./ch05-error-handling/05-errors-is-as
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type PaymentError struct {
	Code   string
	Amount float64
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("payment %s for Rs %.2f", e.Code, e.Amount)
}

func findUser(id int) error {
	if id != 1 {
		return fmt.Errorf("find user %d: %w", id, ErrNotFound)
	}
	return nil
}

func charge(amount float64) error {
	if amount > 10000 {
		return fmt.Errorf("checkout: %w", &PaymentError{"DECLINED_LIMIT", amount})
	}
	return nil
}

func main() {
	// errors.Is: "is this error (or anything it wraps) that specific value?"
	err := findUser(42)
	fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("-> show a 'user not found' page")
	}
	// == would fail because err is wrapped:
	fmt.Println("err == ErrNotFound?", err == ErrNotFound)

	// errors.As: "is there an error of this TYPE inside? give it to me"
	err = charge(25000)
	var pe *PaymentError
	if errors.As(err, &pe) {
		fmt.Println("-> code:", pe.Code, "amount:", pe.Amount)
	}

	// errors.Join combines several errors into one
	all := errors.Join(findUser(7), charge(50000))
	fmt.Println("joined:\n" + all.Error())
	fmt.Println("joined contains ErrNotFound?", errors.Is(all, ErrNotFound))
}
