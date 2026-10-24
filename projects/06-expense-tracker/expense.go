package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
	"time"
)

const dateLayout = "2006-01-02" // Go's reference date, written the way we want ours

// Expense is one line in the tracker. The tags control the JSON file format.
type Expense struct {
	ID       int     `json:"id"`
	Date     string  `json:"date"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Note     string  `json:"note,omitempty"`
}

// NewExpense validates the input before anything is saved.
func NewExpense(date string, amount float64, category, note string) (Expense, error) {
	if _, err := time.Parse(dateLayout, date); err != nil {
		return Expense{}, fmt.Errorf("date %q must look like 2026-10-24", date)
	}
	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be positive, got %.2f", amount)
	}
	category = strings.ToLower(strings.TrimSpace(category))
	if category == "" {
		return Expense{}, fmt.Errorf("category is required")
	}
	return Expense{Date: date, Amount: amount, Category: category, Note: note}, nil
}

// Total is one row of a summary.
type Total struct {
	Key   string
	Sum   float64
	Count int
}

// SummaryBy groups expenses by a key (category or month) and sorts biggest first.
func SummaryBy(list []Expense, key func(Expense) string) []Total {
	groups := map[string]*Total{}
	for _, e := range list {
		k := key(e)
		if groups[k] == nil {
			groups[k] = &Total{Key: k}
		}
		groups[k].Sum += e.Amount
		groups[k].Count++
	}
	totals := make([]Total, 0, len(groups))
	for _, k := range slices.Sorted(maps.Keys(groups)) {
		totals = append(totals, *groups[k])
	}
	slices.SortStableFunc(totals, func(a, b Total) int { return cmp.Compare(b.Sum, a.Sum) })
	return totals
}

func ByCategory(e Expense) string { return e.Category }
func ByMonth(e Expense) string    { return e.Date[:7] } // "2026-10"

// InMonth keeps only expenses from one month ("2026-10"); empty means all.
func InMonth(list []Expense, month string) []Expense {
	if month == "" {
		return list
	}
	var out []Expense
	for _, e := range list {
		if strings.HasPrefix(e.Date, month) {
			out = append(out, e)
		}
	}
	return out
}
