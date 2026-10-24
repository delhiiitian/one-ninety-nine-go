package main

import (
	"fmt"
	"time"
)

const (
	MaxPINTries = 3
	DailyLimit  = 20000.0
)

// Account is the bank's record for one card.
type Account struct {
	Card         string
	Holder       string
	pin          string // lowercase: other packages could not read it
	Balance      float64
	WithdrawnDay float64
	WrongTries   int
	Blocked      bool
	History      []string
}

func (a *Account) record(format string, args ...any) {
	stamp := time.Date(2026, 10, 24, 10, len(a.History), 0, 0, time.UTC).Format("02 Jan 15:04")
	a.History = append(a.History, stamp+"  "+fmt.Sprintf(format, args...))
}

// Bank holds every account, looked up by card number.
type Bank struct {
	accounts map[string]*Account
}

func NewBank() *Bank {
	return &Bank{accounts: map[string]*Account{}}
}

func (b *Bank) Open(card, holder, pin string, balance float64) {
	b.accounts[card] = &Account{Card: card, Holder: holder, pin: pin, Balance: balance}
}

// Login checks the PIN. Three wrong tries block the card.
func (b *Bank) Login(card, pin string) (*Account, error) {
	acc, ok := b.accounts[card]
	if !ok {
		return nil, ErrCardNotFound
	}
	if acc.Blocked {
		return nil, ErrCardBlocked
	}
	if pin != acc.pin {
		acc.WrongTries++
		if acc.WrongTries >= MaxPINTries {
			acc.Blocked = true
			return nil, ErrCardBlocked
		}
		return nil, fmt.Errorf("%w (tries left: %d)", ErrWrongPIN, MaxPINTries-acc.WrongTries)
	}
	acc.WrongTries = 0
	return acc, nil
}

// CanWithdraw checks the bank's rules without moving any money.
func (b *Bank) CanWithdraw(acc *Account, amount float64) error {
	switch {
	case amount <= 0 || int(amount)%100 != 0:
		return ErrNotMultiple
	case amount > acc.Balance:
		return &InsufficientFundsError{acc.Balance, amount}
	case acc.WithdrawnDay+amount > DailyLimit:
		return &LimitError{DailyLimit, DailyLimit - acc.WithdrawnDay}
	}
	return nil
}

// Withdraw applies the rules and then takes the money from the account.
// The ATM checks its own cash separately.
func (b *Bank) Withdraw(acc *Account, amount float64) error {
	if err := b.CanWithdraw(acc, amount); err != nil {
		return err
	}
	acc.Balance -= amount
	acc.WithdrawnDay += amount
	acc.record("withdraw  -%9.2f  bal %9.2f", amount, acc.Balance)
	return nil
}

func (b *Bank) Deposit(acc *Account, amount float64) error {
	if amount <= 0 || int(amount)%100 != 0 {
		return ErrNotMultiple
	}
	acc.Balance += amount
	acc.record("deposit   +%9.2f  bal %9.2f", amount, acc.Balance)
	return nil
}
