// Chapter 4.2: Slices (growable lists)
// Run: go run ./ch04-data-structures/02-slices
package main

import "fmt"

func main() {
	// A slice literal: like an array without a size.
	queue := []string{"Asha", "Bilal"}
	fmt.Println(queue, "len:", len(queue))

	// append adds to the end and returns the new slice.
	queue = append(queue, "Chen")
	queue = append(queue, "Dev", "Esha")
	fmt.Println(queue, "len:", len(queue))

	// Serve the first person: slice off the front.
	served := queue[0]
	queue = queue[1:]
	fmt.Println("served:", served, "remaining:", queue)

	// make creates an empty slice with room to grow (capacity).
	scores := make([]int, 0, 10)
	for i := 1; i <= 5; i++ {
		scores = append(scores, i*10)
	}
	fmt.Println("scores:", scores, "len:", len(scores), "cap:", cap(scores))

	// The zero value of a slice is nil, and append works on it.
	var tags []string
	fmt.Println("nil slice?", tags == nil, "len:", len(tags))
	tags = append(tags, "new", "sale")
	fmt.Println("tags:", tags)

	// A slice of slices: a playlist of albums
	albums := [][]string{
		{"Song A", "Song B"},
		{"Song C"},
	}
	albums[1] = append(albums[1], "Song D")
	fmt.Println("albums:", albums)
}
