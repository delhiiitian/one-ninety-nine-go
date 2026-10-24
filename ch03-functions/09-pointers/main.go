// Chapter 3.9: Pass by value and pointers
// Run: go run ./ch03-functions/09-pointers
package main

import "fmt"

// Go copies arguments. This changes only the copy.
func addBonusCopy(salary int) {
	salary += 5000
}

// A pointer (*int) is the address of the original. Changing *salary changes it.
func addBonus(salary *int) {
	*salary += 5000
}

type Account struct {
	Owner   string
	Balance float64
}

// Pointers to structs avoid copying and let the function update the original.
func deposit(acc *Account, amount float64) {
	acc.Balance += amount // Go lets you write acc.Balance instead of (*acc).Balance
}

func main() {
	salary := 50000
	addBonusCopy(salary)
	fmt.Println("after copy version:", salary) // unchanged

	addBonus(&salary) // & means "address of"
	fmt.Println("after pointer version:", salary)

	p := &salary
	fmt.Println("p points to salary, so *p =", *p)

	acc := Account{"Neha", 1000}
	deposit(&acc, 250)
	fmt.Printf("%s has Rs %.2f\n", acc.Owner, acc.Balance)

	// A pointer that points nowhere is nil. Always check before using.
	var nothing *Account
	if nothing == nil {
		fmt.Println("no account yet")
	}
}
