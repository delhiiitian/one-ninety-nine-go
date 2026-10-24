// Chapter 7.5: gofmt and go vet (the built-in code reviewers)
// Run: go run ./ch07-modules-tooling/05-gofmt-vet
//
// Try the tools on the deliberately messy file next to this one:
//
//	cp ch07-modules-tooling/05-gofmt-vet/messy.go.txt /tmp/messy.go
//	gofmt -d /tmp/messy.go     # shows what gofmt would change
//	gofmt -w /tmp/messy.go     # rewrites the file neatly
//	go vet /tmp/messy.go       # finds real bugs
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
	p := Point{3, 4}
	// go vet checks that Printf verbs match the arguments.
	fmt.Printf("point %d,%d\n", p.X, p.Y)
	fmt.Println("This file is already gofmt-formatted and vet-clean.")
}
