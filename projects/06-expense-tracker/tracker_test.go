package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestSaveAndLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data", "expenses.json")
	s, err := Load(path) // missing file: empty store
	if err != nil {
		t.Fatal(err)
	}
	e, _ := NewExpense("2026-10-02", 450, "Food", "lunch")
	s.Add(e)
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}

	again, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(again.Expenses) != 1 || again.Expenses[0].Category != "food" || again.NextID != 2 {
		t.Fatalf("round trip lost data: %+v", again)
	}
}

func TestValidation(t *testing.T) {
	if _, err := NewExpense("24/10/2026", 10, "food", ""); err == nil {
		t.Error("wrong date format accepted")
	}
	if _, err := NewExpense("2026-10-24", -5, "food", ""); err == nil {
		t.Error("negative amount accepted")
	}
}

func TestSummaryAndCSV(t *testing.T) {
	list := []Expense{
		{1, "2026-09-30", 150, "travel", ""},
		{2, "2026-10-01", 300, "food", ""},
		{3, "2026-10-05", 200, "travel", "auto, metro"},
	}
	got := SummaryBy(list, ByCategory)
	if got[0].Key != "travel" || got[0].Sum != 350 || got[1].Key != "food" {
		t.Fatalf("summary = %+v", got)
	}
	if n := len(InMonth(list, "2026-10")); n != 2 {
		t.Fatalf("InMonth found %d, want 2", n)
	}
	var buf bytes.Buffer
	WriteCSV(&buf, list)
	if !strings.Contains(buf.String(), `"auto, metro"`) {
		t.Fatalf("CSV did not quote a comma:\n%s", buf.String())
	}
}
