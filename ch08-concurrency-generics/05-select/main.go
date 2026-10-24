// Chapter 8.5: select (wait on several channels at once)
// Run: go run ./ch08-concurrency-generics/05-select
package main

import (
	"fmt"
	"time"
)

// quote simulates a price service that answers after a delay.
func quote(company string, price int, delay time.Duration) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		ch <- fmt.Sprintf("%s: Rs %d", company, price)
	}()
	return ch
}

func main() {
	// First answer wins: ask two cab companies, take whoever replies first.
	a := quote("CabA", 240, 150*time.Millisecond)
	b := quote("CabB", 220, 80*time.Millisecond)
	select {
	case q := <-a:
		fmt.Println("first quote", q)
	case q := <-b:
		fmt.Println("first quote", q)
	}

	// Timeout: give up if no answer in 100ms.
	slow := quote("SlowCab", 200, 500*time.Millisecond)
	select {
	case q := <-slow:
		fmt.Println(q)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("SlowCab took too long, skipping")
	}

	// Non-blocking check with default
	msgs := make(chan string, 1)
	select {
	case m := <-msgs:
		fmt.Println("got", m)
	default:
		fmt.Println("no new messages, carrying on")
	}

	// A ticker loop that stops on a signal
	tick := time.NewTicker(50 * time.Millisecond)
	defer tick.Stop()
	stop := time.After(180 * time.Millisecond)
	for n := 1; ; n++ {
		select {
		case <-tick.C:
			fmt.Println("heartbeat", n)
		case <-stop:
			fmt.Println("stopping heartbeat")
			return
		}
	}
}
