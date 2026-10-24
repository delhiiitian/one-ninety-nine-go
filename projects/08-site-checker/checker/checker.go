// Package checker tests many websites at the same time using a pool of workers.
package checker

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result is what we learned about one URL.
type Result struct {
	URL     string
	Status  int           // HTTP status code, 0 if the request failed
	Latency time.Duration // how long the answer took
	Err     error
}

// Up reports whether the site answered with a 2xx or 3xx status.
func (r Result) Up() bool { return r.Err == nil && r.Status >= 200 && r.Status < 400 }

// Checker holds the settings shared by every check.
type Checker struct {
	Client  *http.Client
	Workers int
	Timeout time.Duration // per request
}

// New returns a Checker with sensible defaults.
func New(workers int, timeout time.Duration) *Checker {
	return &Checker{
		Client:  &http.Client{},
		Workers: max(workers, 1),
		Timeout: timeout,
	}
}

// Check requests one URL, giving up after c.Timeout or when ctx is cancelled.
func (c *Checker) Check(ctx context.Context, url string) Result {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Result{URL: url, Err: fmt.Errorf("bad URL: %w", err)}
	}
	req.Header.Set("User-Agent", "go199-site-checker/1.0")

	resp, err := c.Client.Do(req)
	latency := time.Since(start)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			err = fmt.Errorf("timed out after %v", c.Timeout)
		}
		return Result{URL: url, Latency: latency, Err: err}
	}
	resp.Body.Close()
	return Result{URL: url, Status: resp.StatusCode, Latency: latency}
}

// CheckAll checks every URL using c.Workers goroutines.
// Results come back in the same order as the input, whatever order they finish in.
func (c *Checker) CheckAll(ctx context.Context, urls []string) []Result {
	results := make([]Result, len(urls))
	jobs := make(chan int) // indexes into urls

	var wg sync.WaitGroup
	for range c.Workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				results[i] = c.Check(ctx, urls[i]) // each worker writes its own slot: no lock needed
			}
		}()
	}

	for i := range urls {
		select {
		case jobs <- i:
		case <-ctx.Done(): // stop handing out work if the whole run is cancelled
		}
	}
	close(jobs)
	wg.Wait()

	for i, r := range results {
		if r.URL == "" { // never started because ctx was cancelled
			results[i] = Result{URL: urls[i], Err: ctx.Err()}
		}
	}
	return results
}
