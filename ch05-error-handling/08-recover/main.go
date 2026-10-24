// Chapter 5.8: recover (catch a panic at a safe boundary)
// Run: go run ./ch05-error-handling/08-recover
package main

import (
	"errors"
	"fmt"
)

// safely runs a job and turns any panic into an ordinary error.
// Web servers do this for every request so one bad request cannot crash the site.
func safely(name string, job func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("job %q crashed: %v", name, r)
		}
	}()
	job()
	return nil
}

func divide(a, b int) int { return a / b }

func main() {
	jobs := []struct {
		name string
		run  func()
	}{
		{"report", func() { fmt.Println("  report generated") }},
		{"average", func() { fmt.Println("  average:", divide(100, 0)) }},
		{"lookup", func() {
			var prices map[string]int
			prices["tea"] = 20 // writing to a nil map panics
		}},
		{"email", func() { fmt.Println("  email sent") }},
	}

	var failures []error
	for _, j := range jobs {
		fmt.Println("running", j.name)
		if err := safely(j.name, j.run); err != nil {
			fmt.Println("  !", err)
			failures = append(failures, err)
		}
	}
	fmt.Printf("\n%d jobs ran, %d failed, program still alive.\n", len(jobs), len(failures))
	fmt.Println(errors.Join(failures...))
}
