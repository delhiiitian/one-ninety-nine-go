// Chapter 8.2: sync.WaitGroup (wait for everyone to finish)
// Run: go run ./ch08-concurrency-generics/02-waitgroup
package main

import (
	"fmt"
	"sync"
	"time"
)

func download(file string, sizeMB int) {
	time.Sleep(time.Duration(sizeMB) * 20 * time.Millisecond) // pretend network time
	fmt.Printf("downloaded %-12s (%d MB)\n", file, sizeMB)
}

func main() {
	files := map[string]int{"photo.jpg": 3, "video.mp4": 10, "notes.pdf": 1, "song.mp3": 5}

	start := time.Now()
	var wg sync.WaitGroup
	for name, size := range files {
		wg.Add(1) // one more job to wait for
		go func() {
			defer wg.Done() // this job is finished
			download(name, size)
		}()
	}
	wg.Wait() // block until the counter is back to zero
	fmt.Printf("all %d downloads finished in about %v (the slowest one)\n",
		len(files), time.Since(start).Round(100*time.Millisecond))
}
