// Chapter 5.9: Validation and guard clauses
// Run: go run ./ch05-error-handling/09-validation
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Transfer struct {
	From, To string
	Amount   float64
	Note     string
}

// Validate collects every problem at once, so the user can fix them all in one go.
func (t Transfer) Validate() error {
	var errs []error
	if strings.TrimSpace(t.From) == "" {
		errs = append(errs, errors.New("from account is required"))
	}
	if strings.TrimSpace(t.To) == "" {
		errs = append(errs, errors.New("to account is required"))
	}
	if t.From != "" && t.From == t.To {
		errs = append(errs, errors.New("cannot transfer to the same account"))
	}
	if t.Amount <= 0 {
		errs = append(errs, fmt.Errorf("amount must be positive, got %.2f", t.Amount))
	}
	if t.Amount > 200000 {
		errs = append(errs, errors.New("amount above daily limit of 2,00,000"))
	}
	if len(t.Note) > 50 {
		errs = append(errs, errors.New("note must be 50 characters or fewer"))
	}
	return errors.Join(errs...) // nil when errs is empty
}

// Process uses guard clauses: handle problems first, keep the happy path flat.
func Process(t Transfer, balances map[string]float64) error {
	if err := t.Validate(); err != nil {
		return fmt.Errorf("invalid transfer:\n%w", err)
	}
	if _, ok := balances[t.From]; !ok {
		return fmt.Errorf("unknown account %s", t.From)
	}
	if balances[t.From] < t.Amount {
		return fmt.Errorf("insufficient balance in %s", t.From)
	}

	// Happy path: no nesting needed
	balances[t.From] -= t.Amount
	balances[t.To] += t.Amount
	return nil
}

func main() {
	balances := map[string]float64{"ACC1": 5000, "ACC2": 100}

	transfers := []Transfer{
		{"ACC1", "ACC2", 1500, "rent share"},
		{"", "ACC1", -5, ""},
		{"ACC2", "ACC2", 300000, ""},
		{"ACC2", "ACC1", 999, "dinner"},
	}
	for i, t := range transfers {
		if err := Process(t, balances); err != nil {
			fmt.Printf("transfer %d failed: %v\n\n", i+1, err)
			continue
		}
		fmt.Printf("transfer %d ok: %v\n\n", i+1, balances)
	}
}
