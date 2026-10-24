// Project 8: Concurrent Website Checker
//
//	go run ./projects/08-site-checker -demo                       # local demo site, no internet needed
//	go run ./projects/08-site-checker https://go.dev https://example.com
//	go run ./projects/08-site-checker -file urls.txt -workers 20 -timeout 3s
//	go test ./projects/08-site-checker/...
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/delhiiitian/one-ninety-nine-go/projects/08-site-checker/checker"
)

func main() {
	workers := flag.Int("workers", 10, "how many sites to check at the same time")
	timeout := flag.Duration("timeout", 2*time.Second, "give up on a site after this long")
	file := flag.String("file", "", "read URLs from this file, one per line")
	demo := flag.Bool("demo", false, "check a built-in local demo site")
	flag.Parse()

	urls, err := collectURLs(*file, flag.Args(), *demo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "no URLs given. Try: -demo, or pass URLs, or -file urls.txt")
		os.Exit(2)
	}

	// Ctrl+C cancels the context, so every in-flight request stops cleanly.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c := checker.New(*workers, *timeout)
	start := time.Now()
	results := c.CheckAll(ctx, urls)
	printReport(results, time.Since(start), *workers)

	for _, r := range results {
		if !r.Up() {
			os.Exit(1) // non-zero exit lets scripts and CI notice outages
		}
	}
}

func collectURLs(file string, args []string, demo bool) ([]string, error) {
	urls := append([]string{}, args...)
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if line != "" && !strings.HasPrefix(line, "#") {
				urls = append(urls, line)
			}
		}
		if err := sc.Err(); err != nil {
			return nil, err
		}
	}
	if demo {
		base, err := startDemoServer()
		if err != nil {
			return nil, err
		}
		for _, p := range []string{"/", "/shop", "/search", "/checkout", "/old-page", "/missing"} {
			urls = append(urls, base+p)
		}
	}
	return urls, nil
}

func printReport(results []checker.Result, took time.Duration, workers int) {
	up := 0
	fmt.Printf("%-6s %-6s %8s  %s\n", "STATE", "CODE", "TIME", "URL")
	for _, r := range results {
		state, code := "DOWN", "-"
		if r.Up() {
			state = "UP"
			up++
		}
		if r.Status != 0 {
			code = fmt.Sprint(r.Status)
		}
		line := fmt.Sprintf("%-6s %-6s %8v  %s", state, code, r.Latency.Round(time.Millisecond), r.URL)
		if r.Err != nil {
			line += "  (" + shorten(r.Err.Error()) + ")"
		}
		fmt.Println(line)
	}
	fmt.Printf("\n%d of %d up. Checked in %v with %d workers.\n", up, len(results), took.Round(time.Millisecond), workers)
}

func shorten(s string) string {
	if i := strings.LastIndex(s, ": "); i >= 0 && len(s) > 60 {
		return s[i+2:]
	}
	return s
}
