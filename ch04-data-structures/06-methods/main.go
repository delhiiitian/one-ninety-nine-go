// Chapter 4.6: Methods (functions that belong to a type)
// Run: go run ./ch04-data-structures/06-methods
package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Width, Height float64
}

// A value receiver: reads the rectangle, does not change it.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Wallet struct {
	Owner   string
	Balance float64
	History []string
}

// A pointer receiver: changes the wallet itself.
func (w *Wallet) Add(amount float64) {
	w.Balance += amount
	w.History = append(w.History, fmt.Sprintf("+%.2f", amount))
}

func (w *Wallet) Pay(amount float64) error {
	if amount > w.Balance {
		return errors.New("not enough money")
	}
	w.Balance -= amount
	w.History = append(w.History, fmt.Sprintf("-%.2f", amount))
	return nil
}

// String is special: fmt uses it automatically when printing.
func (w Wallet) String() string {
	return fmt.Sprintf("%s's wallet: Rs %.2f", w.Owner, w.Balance)
}

// Methods work on any type you define, not only structs.
type Celsius float64

func (c Celsius) Fahrenheit() float64 { return float64(c)*9/5 + 32 }

func main() {
	room := Rectangle{4, 5}
	fmt.Println("area:", room.Area())

	w := Wallet{Owner: "Farah"}
	w.Add(500) // Go takes &w for you
	if err := w.Pay(120); err != nil {
		fmt.Println(err)
	}
	if err := w.Pay(1000); err != nil {
		fmt.Println("payment failed:", err)
	}
	fmt.Println(w)
	fmt.Println("history:", w.History)

	fmt.Println("37C =", Celsius(37).Fahrenheit(), "F")
}
