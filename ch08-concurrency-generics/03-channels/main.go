// Chapter 8.3: Channels (passing results between goroutines)
// Run: go run ./ch08-concurrency-generics/03-channels
package main

import (
	"fmt"
	"strings"
)

type Order struct {
	ID   int
	Dish string
}

// kitchen receives orders on one channel and sends finished plates on another.
// The arrows in the types say which way data may flow.
func kitchen(orders <-chan Order, plates chan<- string) {
	for o := range orders { // loops until the orders channel is closed
		plates <- fmt.Sprintf("order %d: %s ready", o.ID, strings.ToUpper(o.Dish))
	}
	close(plates) // tell the receiver there is nothing more
}

func main() {
	orders := make(chan Order)
	plates := make(chan string)

	go kitchen(orders, plates)

	// The waiter sends orders in its own goroutine.
	go func() {
		for i, dish := range []string{"dosa", "idli", "vada"} {
			orders <- Order{i + 1, dish} // blocks until the kitchen takes it
		}
		close(orders)
	}()

	// main receives plates until the channel is closed.
	for p := range plates {
		fmt.Println(p)
	}

	// Receiving from a closed channel gives the zero value and ok=false.
	v, ok := <-plates
	fmt.Printf("after close: %q ok=%v\n", v, ok)
}
