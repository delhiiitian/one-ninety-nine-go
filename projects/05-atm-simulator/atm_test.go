package main

import (
	"errors"
	"testing"
)

func TestLoginBlocksAfterThreeWrongPINs(t *testing.T) {
	b := NewBank()
	b.Open("1", "A", "1234", 100)
	for i := 0; i < 2; i++ {
		if _, err := b.Login("1", "0000"); !errors.Is(err, ErrWrongPIN) {
			t.Fatalf("try %d: want ErrWrongPIN, got %v", i+1, err)
		}
	}
	if _, err := b.Login("1", "0000"); !errors.Is(err, ErrCardBlocked) {
		t.Fatalf("want ErrCardBlocked, got %v", err)
	}
	if _, err := b.Login("1", "1234"); !errors.Is(err, ErrCardBlocked) {
		t.Fatal("a blocked card must stay blocked even with the right PIN")
	}
}

func TestWithdrawRules(t *testing.T) {
	b := NewBank()
	b.Open("1", "A", "1", 30000)
	acc, _ := b.Login("1", "1")

	if err := b.Withdraw(acc, 250); !errors.Is(err, ErrNotMultiple) {
		t.Fatalf("want ErrNotMultiple, got %v", err)
	}
	var funds *InsufficientFundsError
	if err := b.Withdraw(acc, 50000); !errors.As(err, &funds) || funds.Balance != 30000 {
		t.Fatalf("want InsufficientFundsError, got %v", err)
	}
	if err := b.Withdraw(acc, 15000); err != nil {
		t.Fatal(err)
	}
	var limit *LimitError
	if err := b.Withdraw(acc, 6000); !errors.As(err, &limit) || limit.Remaining != 5000 {
		t.Fatalf("want LimitError with 5000 left, got %v", err)
	}
}

func TestATMNotes(t *testing.T) {
	m := NewATM(map[int]int{500: 1, 200: 2, 100: 0})
	plan, err := m.Plan(900)
	if err != nil || plan[500] != 1 || plan[200] != 2 {
		t.Fatalf("plan(900) = %v, %v", plan, err)
	}
	if _, err := m.Plan(1000); !errors.Is(err, ErrNoCashInATM) {
		t.Fatalf("want ErrNoCashInATM, got %v", err)
	}
}

func TestSafelyRecovers(t *testing.T) {
	err := safely(func() error { panic("boom") })
	if err == nil {
		t.Fatal("expected the panic to become an error")
	}
}
