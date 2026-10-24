package main

import "testing"

func TestCalculate(t *testing.T) {
	b := Calculate(1000, 10, 4)
	if b.Total != 1150 {
		t.Fatalf("total = %.2f, want 1150", b.Total)
	}
	if b.PerPerson != 287.5 {
		t.Fatalf("per person = %.2f, want 287.50", b.PerPerson)
	}
}

func TestRoundUpNeverUnderpays(t *testing.T) {
	b := Calculate(100, 0, 3) // 105 / 3 = 35 exactly
	if b.PerPerson*3 < b.Total {
		t.Fatalf("group pays %.2f, less than %.2f", b.PerPerson*3, b.Total)
	}
	b = Calculate(101, 0, 3) // 106.05 / 3 = 35.35
	if b.PerPerson*3 < b.Total {
		t.Fatalf("group pays %.2f, less than %.2f", b.PerPerson*3, b.Total)
	}
}
