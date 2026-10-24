// Chapter 8.8: context (cancel and time out work)
// Run: go run ./ch08-concurrency-generics/08-context
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// searchFlights pretends to call an airline. It stops early if ctx is cancelled.
func searchFlights(ctx context.Context, airline string, takes time.Duration) (string, error) {
	select {
	case <-time.After(takes):
		return airline + ": 3 flights found", nil
	case <-ctx.Done(): // cancelled or timed out
		return "", fmt.Errorf("%s: %w", airline, ctx.Err())
	}
}

func main() {
	// 1. A timeout for the whole search
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel() // always release the context's resources

	for _, a := range []struct {
		name  string
		takes time.Duration
	}{{"IndiGo", 50 * time.Millisecond}, {"AirSlow", 400 * time.Millisecond}} {
		res, err := searchFlights(ctx, a.name, a.takes)
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("gave up:", err)
			continue
		}
		fmt.Println(res)
	}

	// 2. Manual cancel: the user closes the page, so stop the background work.
	ctx2, cancel2 := context.WithCancel(context.Background())
	progress := make(chan int)
	go func() {
		defer close(progress)
		for pct := 10; pct <= 100; pct += 10 {
			select {
			case <-ctx2.Done():
				return
			case <-time.After(20 * time.Millisecond):
			}
			select {
			case progress <- pct:
			case <-ctx2.Done():
				return
			}
		}
	}()
	for pct := range progress {
		fmt.Printf("report %d%% ", pct)
		if pct == 40 {
			fmt.Println("\nuser closed the page: cancelling")
			cancel2()
		}
	}
	fmt.Println("background work stopped:", ctx2.Err())
}
