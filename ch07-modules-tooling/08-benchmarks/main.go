// Chapter 7.8: Benchmarks and examples
// Run the program:    go run ./ch07-modules-tooling/08-benchmarks
// Run the benchmarks: go test -bench . -benchmem ./ch07-modules-tooling/08-benchmarks
package main

import (
	"fmt"
	"time"
)

func main() {
	list := make([]string, 5000)
	for i := range list {
		list[i] = "word"
	}
	for _, f := range []struct {
		name string
		fn   func([]string) string
	}{{"JoinPlus", JoinPlus}, {"JoinBuilder", JoinBuilder}} {
		start := time.Now()
		f.fn(list)
		fmt.Printf("%-12s took %v\n", f.name, time.Since(start).Round(time.Microsecond))
	}
	fmt.Println("For reliable numbers use: go test -bench . -benchmem ./ch07-modules-tooling/08-benchmarks")
}
