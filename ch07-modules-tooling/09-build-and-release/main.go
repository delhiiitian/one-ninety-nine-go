// Chapter 7.9: Building, cross-compiling and version stamps
//
// Run:            go run ./ch07-modules-tooling/09-build-and-release
// Build a binary: go build -o hello ./ch07-modules-tooling/09-build-and-release && ./hello
// Stamp version:  go build -ldflags "-X main.version=1.2.0" -o hello ./ch07-modules-tooling/09-build-and-release
// For Windows:    GOOS=windows GOARCH=amd64 go build -o hello.exe ./ch07-modules-tooling/09-build-and-release
// For a Pi:       GOOS=linux GOARCH=arm64 go build -o hello-pi ./ch07-modules-tooling/09-build-and-release
// Install to PATH: go install ./ch07-modules-tooling/09-build-and-release
package main

import (
	"fmt"
	"runtime"
)

// version is replaced at build time with -ldflags "-X main.version=..."
var version = "dev"

func main() {
	fmt.Println("hello tool, version", version)
	fmt.Printf("running on %s/%s, built with %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
	fmt.Println("CPU cores available:", runtime.NumCPU())
}
