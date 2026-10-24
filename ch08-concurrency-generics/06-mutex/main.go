// Chapter 8.6: sync.Mutex (one at a time, please)
// Run: go run ./ch08-concurrency-generics/06-mutex
// Try:  go run -race ./ch08-concurrency-generics/06-mutex  (needs a C compiler on some systems)
package main

import (
	"fmt"
	"sync"
)

// UnsafeCounter has a data race: many goroutines change n at the same time.
type UnsafeCounter struct{ n int }

func (c *UnsafeCounter) Inc() { c.n++ }

// SafeCounter guards n with a mutex, like a key to a single bathroom.
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// Ticket booking: check-then-act must happen under one lock.
type Show struct {
	mu   sync.Mutex
	left int
}

func (s *Show) Book() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.left == 0 {
		return false
	}
	s.left--
	return true
}

func run(times int, f func()) {
	var wg sync.WaitGroup
	for range times {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}
	wg.Wait()
}

func main() {
	var bad UnsafeCounter
	run(10_000, bad.Inc)
	fmt.Println("unsafe counter (should be 10000):", bad.n, "<- often wrong!")

	var good SafeCounter
	run(10_000, good.Inc)
	fmt.Println("safe counter:", good.Value())

	show := &Show{left: 100}
	var booked SafeCounter
	run(500, func() {
		if show.Book() {
			booked.Inc()
		}
	})
	fmt.Println("500 people tried, tickets sold:", booked.Value(), "left:", show.left)
}
