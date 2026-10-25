// Chapter 9.1: Time, dates and durations
// Run: go run ./ch09-real-world/01-time
package main

import (
	"fmt"
	"time"
	_ "time/tzdata" // embed the time-zone database so this works on every machine
)

func main() {
	// A fixed moment so the output is the same every run.
	// In real code use time.Now().
	boarding := time.Date(2026, time.October, 24, 18, 45, 0, 0, time.UTC)

	// Layouts are written using Go's reference time: Mon Jan 2 15:04:05 2006
	fmt.Println("ISO:      ", boarding.Format(time.RFC3339))
	fmt.Println("friendly: ", boarding.Format("Mon, 2 Jan 2006 at 3:04 PM"))
	fmt.Println("date only:", boarding.Format("02/01/2006"))

	// Time zones
	ist, _ := time.LoadLocation("Asia/Kolkata")
	ny, _ := time.LoadLocation("America/New_York")
	fmt.Println("in India: ", boarding.In(ist).Format("15:04 MST"))
	fmt.Println("in NY:    ", boarding.In(ny).Format("15:04 MST"))

	// Durations and arithmetic
	flight := 9*time.Hour + 30*time.Minute
	landing := boarding.Add(flight)
	fmt.Println("flight time:", flight, "lands:", landing.In(ny).Format("Mon 15:04 MST"))

	due := boarding.AddDate(0, 0, 30) // 30 days later
	fmt.Println("invoice due:", due.Format("2 Jan 2006"))

	// Difference between two times
	birthday := time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)
	left := birthday.Sub(boarding)
	fmt.Printf("days until 25 Dec: %.0f\n", left.Hours()/24)

	// Parsing text into a time
	t, err := time.Parse("2006-01-02 15:04", "2026-11-14 09:30")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("parsed:", t.Format(time.Kitchen), t.Weekday())

	// Comparing
	fmt.Println("parsed is after boarding?", t.After(boarding))

	// Truncate and round: group events by hour
	fmt.Println("start of hour:", boarding.Add(17*time.Minute).Truncate(time.Hour).Format("15:04"))

	// Measuring how long code takes
	start := time.Now()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("took roughly", time.Since(start).Round(10*time.Millisecond))
}
