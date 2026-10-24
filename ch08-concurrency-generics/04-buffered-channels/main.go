// Chapter 8.4: Buffered channels (a queue with limited space)
// Run: go run ./ch08-concurrency-generics/04-buffered-channels
package main

import (
	"fmt"
	"time"
)

func main() {
	// A conveyor belt that holds 3 items.
	belt := make(chan string, 3)

	belt <- "box 1" // these do not block: there is room
	belt <- "box 2"
	belt <- "box 3"
	fmt.Println("on the belt:", len(belt), "of", cap(belt))
	// belt <- "box 4" // would block forever here: the belt is full

	fmt.Println("took", <-belt)
	belt <- "box 4" // room again
	close(belt)
	for b := range belt {
		fmt.Println("took", b)
	}

	// A buffered channel as a limiter: at most 2 "printers" busy at a time.
	printers := make(chan struct{}, 2)
	done := make(chan string)
	start := time.Now()
	for _, job := range []string{"invoice", "ticket", "report", "label"} {
		go func() {
			printers <- struct{}{}        // take a printer (waits if both are busy)
			defer func() { <-printers }() // give it back
			time.Sleep(100 * time.Millisecond)
			done <- job
		}()
	}
	for range 4 {
		<-done
	}
	fmt.Printf("4 jobs on 2 printers took about %v\n", time.Since(start).Round(100*time.Millisecond))
}
