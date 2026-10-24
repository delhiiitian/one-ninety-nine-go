package main

import (
	"fmt"
	"slices"
)

// ATM holds a limited stock of notes, like a real machine.
type ATM struct {
	notes map[int]int // note value -> how many are loaded
}

func NewATM(notes map[int]int) *ATM { return &ATM{notes: notes} }

// Plan works out which notes to give, largest first, without changing the stock.
func (m *ATM) Plan(amount int) (map[int]int, error) {
	values := make([]int, 0, len(m.notes))
	for v := range m.notes {
		values = append(values, v)
	}
	slices.Sort(values)
	slices.Reverse(values)

	plan := map[int]int{}
	left := amount
	for _, v := range values {
		use := min(left/v, m.notes[v])
		if use > 0 {
			plan[v] = use
			left -= use * v
		}
	}
	if left != 0 {
		return nil, ErrNoCashInATM
	}
	return plan, nil
}

// Dispense removes the planned notes from the machine.
func (m *ATM) Dispense(plan map[int]int) {
	for v, n := range plan {
		m.notes[v] -= n
	}
}

func describe(plan map[int]int) string {
	s := ""
	for _, v := range []int{2000, 500, 200, 100} {
		if n := plan[v]; n > 0 {
			s += fmt.Sprintf("%d x Rs %d  ", n, v)
		}
	}
	return s
}
