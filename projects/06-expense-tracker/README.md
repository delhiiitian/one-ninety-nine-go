# Project 6: Expense Tracker

Track daily spending from the command line. Expenses are saved to a JSON file so they
survive between runs, summaries show where the money goes by category and by month,
and you can export everything to CSV for Excel or Google Sheets.

Book: *Go in 199 Pages*, Chapter 6 project.

## Run it

```bash
export EXPENSES_FILE=/tmp/my-expenses.json   # optional; default is ./expenses.json
go run ./projects/06-expense-tracker add 2026-10-02 450 food "lunch with team"
go run ./projects/06-expense-tracker add 2026-10-03 1200 travel "train to Mysuru"
go run ./projects/06-expense-tracker list
go run ./projects/06-expense-tracker summary 2026-10
go run ./projects/06-expense-tracker export /tmp/expenses.csv
go test ./projects/06-expense-tracker
```

## Files

| File | Job |
| --- | --- |
| `main.go` | Reads the command and its arguments, prints results |
| `expense.go` | The `Expense` type, validation, month filter and summaries |
| `store.go` | Loads and safely saves the JSON file (write temp file, then rename) |
| `export.go` | Writes CSV to any `io.Writer` |
| `tracker_test.go` | Round-trip, validation, summary and CSV tests |

## Ideas to extend

- Add monthly budgets per category and warn when one is exceeded.
- Add an `import` command that reads a bank statement CSV.
- Show the summary bars in colour.
