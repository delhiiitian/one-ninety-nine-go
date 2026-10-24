// Chapter 8.7: Worker pools (a fixed team sharing a pile of work)
// Run: go run ./ch08-concurrency-generics/07-worker-pool
package main

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Pages int
}

type Result struct {
	JobID    int
	WorkerID int
	Words    int
}

// worker takes jobs until the jobs channel is closed.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(time.Duration(j.Pages) * 10 * time.Millisecond) // "translate" the pages
		results <- Result{j.ID, id, j.Pages * 300}
	}
}

func main() {
	const workers = 3
	jobs := make(chan Job)
	results := make(chan Result)

	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Feed the jobs, then close the channel so workers know to stop.
	go func() {
		for i := 1; i <= 9; i++ {
			jobs <- Job{ID: i, Pages: i%4 + 1}
		}
		close(jobs)
	}()

	// Close results once every worker has finished.
	go func() {
		wg.Wait()
		close(results)
	}()

	start := time.Now()
	var all []Result
	for r := range results {
		all = append(all, r)
	}
	slices.SortFunc(all, func(a, b Result) int { return a.JobID - b.JobID })

	total := 0
	perWorker := map[int]int{}
	for _, r := range all {
		total += r.Words
		perWorker[r.WorkerID]++
	}
	alone := time.Duration(total/300) * 10 * time.Millisecond
	fmt.Println("jobs done:", len(all), "words translated:", total)
	for w := 1; w <= workers; w++ {
		fmt.Printf("worker %d handled %d jobs\n", w, perWorker[w])
	}
	fmt.Printf("time: about %v (one worker alone would need about %v)\n",
		time.Since(start).Round(10*time.Millisecond), alone)
}
