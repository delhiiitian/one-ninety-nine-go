// Project 6: Expense Tracker (saves to a JSON file, exports CSV)
//
//	go run ./projects/06-expense-tracker add 2026-10-02 450 food "lunch with team"
//	go run ./projects/06-expense-tracker list [2026-10]
//	go run ./projects/06-expense-tracker summary [2026-10]
//	go run ./projects/06-expense-tracker delete 3
//	go run ./projects/06-expense-tracker export expenses.csv
//
// Data goes to expenses.json in the current folder, or to $EXPENSES_FILE.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: add | list | summary | delete | export (see the top of main.go)")
	}
	path := os.Getenv("EXPENSES_FILE")
	if path == "" {
		path = "expenses.json"
	}
	store, err := Load(path)
	if err != nil {
		return err
	}

	cmd, rest := args[0], args[1:]
	switch cmd {
	case "add":
		if len(rest) < 3 {
			return fmt.Errorf("usage: add <yyyy-mm-dd> <amount> <category> [note]")
		}
		amount, err := strconv.ParseFloat(rest[1], 64)
		if err != nil {
			return fmt.Errorf("amount %q is not a number", rest[1])
		}
		e, err := NewExpense(rest[0], amount, rest[2], strings.Join(rest[3:], " "))
		if err != nil {
			return err
		}
		e = store.Add(e)
		fmt.Printf("added #%d: %s %.2f %s\n", e.ID, e.Date, e.Amount, e.Category)
		return store.Save()

	case "list":
		list := InMonth(store.Expenses, first(rest))
		if len(list) == 0 {
			fmt.Println("no expenses yet")
			return nil
		}
		total := 0.0
		for _, e := range list {
			fmt.Printf("#%-3d %s  %9.2f  %-10s %s\n", e.ID, e.Date, e.Amount, e.Category, e.Note)
			total += e.Amount
		}
		fmt.Printf("%17s %9.2f\n", "TOTAL", total)

	case "summary":
		list := InMonth(store.Expenses, first(rest))
		grand := 0.0
		for _, e := range list {
			grand += e.Amount
		}
		fmt.Println("By category:")
		for _, t := range SummaryBy(list, ByCategory) {
			fmt.Printf("  %-10s %9.2f  %5.1f%%  %s\n", t.Key, t.Sum, 100*t.Sum/grand, bar(t.Sum/grand))
		}
		fmt.Println("By month:")
		for _, t := range SummaryBy(list, ByMonth) {
			fmt.Printf("  %-10s %9.2f  (%d expenses)\n", t.Key, t.Sum, t.Count)
		}

	case "delete":
		id, err := strconv.Atoi(first(rest))
		if err != nil {
			return fmt.Errorf("usage: delete <id>")
		}
		if err := store.Delete(id); err != nil {
			return err
		}
		fmt.Println("deleted", id)
		return store.Save()

	case "export":
		out := first(rest)
		if out == "" {
			return WriteCSV(os.Stdout, store.Expenses)
		}
		f, err := os.Create(out)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := WriteCSV(f, store.Expenses); err != nil {
			return err
		}
		fmt.Println("exported", len(store.Expenses), "expenses to", out)

	default:
		return fmt.Errorf("unknown command %q", cmd)
	}
	return nil
}

func first(s []string) string {
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

// bar draws a little text chart: 20 blocks = 100%.
func bar(fraction float64) string {
	return strings.Repeat("#", int(fraction*20+0.5))
}
