package main

import (
	"errors"
	"fmt"
)

// Sentinel errors: fixed situations the screen reacts to.
var (
	ErrCardNotFound = errors.New("card not recognised")
	ErrWrongPIN     = errors.New("incorrect PIN")
	ErrCardBlocked  = errors.New("card blocked after too many wrong PINs")
	ErrNotMultiple  = errors.New("amount must be a multiple of 100")
	ErrNoCashInATM  = errors.New("this ATM cannot dispense that amount right now")
)

// InsufficientFundsError carries the numbers so the screen can show them.
type InsufficientFundsError struct {
	Balance, Requested float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: balance Rs %.2f, requested Rs %.2f", e.Balance, e.Requested)
}

// LimitError says how much of today's limit is left.
type LimitError struct {
	Limit, Remaining float64
}

func (e *LimitError) Error() string {
	return fmt.Sprintf("daily limit Rs %.0f reached: Rs %.0f left today", e.Limit, e.Remaining)
}
