// Project 5: ATM Simulator
// Run: go run ./projects/05-atm-simulator
// Demo cards: 1111 (PIN 1234, Ravi), 2222 (PIN 0000, Sara)
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bank := NewBank()
	bank.Open("1111", "Ravi", "1234", 45000)
	bank.Open("2222", "Sara", "0000", 3200)
	atm := NewATM(map[int]int{500: 60, 200: 10, 100: 10})

	in := bufio.NewScanner(os.Stdin)
	ask := func(q string) (string, bool) {
		fmt.Print(q)
		if !in.Scan() {
			return "", false
		}
		return strings.TrimSpace(in.Text()), true
	}

	fmt.Println("=== WELCOME TO GOBANK ATM ===")
	for {
		card, ok := ask("\nInsert card (card number, or 'off'): ")
		if !ok || card == "off" {
			fmt.Println("ATM switched off.")
			return
		}
		acc := login(bank, card, ask)
		if acc == nil {
			continue
		}
		session(bank, atm, acc, ask)
	}
}

func login(bank *Bank, card string, ask func(string) (string, bool)) *Account {
	for {
		pin, ok := ask("Enter PIN: ")
		if !ok {
			return nil
		}
		acc, err := bank.Login(card, pin)
		switch {
		case err == nil:
			fmt.Printf("Hello, %s!\n", acc.Holder)
			return acc
		case errors.Is(err, ErrWrongPIN):
			fmt.Println("!", err)
		default: // card unknown or blocked: give the card back
			fmt.Println("!", err, "- please take your card.")
			return nil
		}
	}
}

func session(bank *Bank, atm *ATM, acc *Account, ask func(string) (string, bool)) {
	for {
		choice, ok := ask("\n1) Balance  2) Withdraw  3) Deposit  4) Mini statement  5) Exit\nChoose: ")
		if !ok || choice == "5" {
			fmt.Println("Thank you. Please take your card.")
			return
		}
		// Each action runs inside a safety net: a bug ends that action, not the ATM.
		if err := safely(func() error { return act(bank, atm, acc, choice, ask) }); err != nil {
			explain(err)
		}
	}
}

func act(bank *Bank, atm *ATM, acc *Account, choice string, ask func(string) (string, bool)) error {
	switch choice {
	case "1":
		fmt.Printf("Balance: Rs %.2f\n", acc.Balance)
	case "2":
		amount, err := askAmount(ask)
		if err != nil {
			return err
		}
		// Check the bank's rules first, then the machine's notes, and only then move money.
		if err := bank.CanWithdraw(acc, amount); err != nil {
			return err
		}
		plan, err := atm.Plan(int(amount))
		if err != nil {
			return err
		}
		if err := bank.Withdraw(acc, amount); err != nil {
			return err
		}
		atm.Dispense(plan)
		fmt.Printf("Please take your cash: %s\nBalance: Rs %.2f\n", describe(plan), acc.Balance)
	case "3":
		amount, err := askAmount(ask)
		if err != nil {
			return err
		}
		if err := bank.Deposit(acc, amount); err != nil {
			return err
		}
		fmt.Printf("Deposited. Balance: Rs %.2f\n", acc.Balance)
	case "4":
		if len(acc.History) == 0 {
			fmt.Println("No transactions yet.")
		}
		start := max(0, len(acc.History)-5)
		for _, line := range acc.History[start:] {
			fmt.Println(line)
		}
	default:
		fmt.Println("Please choose 1 to 5.")
	}
	return nil
}

func askAmount(ask func(string) (string, bool)) (float64, error) {
	text, _ := ask("Amount (Rs): ")
	amount, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("%q is not an amount: %w", text, ErrNotMultiple)
	}
	return amount, nil
}

// explain turns any error into a friendly screen message.
func explain(err error) {
	var funds *InsufficientFundsError
	var limit *LimitError
	switch {
	case errors.As(err, &funds):
		fmt.Printf("! Not enough money. You can withdraw up to Rs %.0f.\n", funds.Balance)
	case errors.As(err, &limit):
		fmt.Printf("! Daily limit reached. You can still take Rs %.0f today.\n", limit.Remaining)
	case errors.Is(err, ErrNotMultiple):
		fmt.Println("! Please enter an amount in hundreds, like 500 or 2300.")
	case errors.Is(err, ErrNoCashInATM):
		fmt.Println("! This ATM is low on notes. Try a smaller amount.")
	default:
		fmt.Println("! Something went wrong:", err)
	}
}

// safely converts a panic into an error so the machine keeps running.
func safely(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("internal error: %v", r)
		}
	}()
	return f()
}
