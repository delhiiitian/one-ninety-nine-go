// Chapter 7.1: Modules (go.mod)
// Run: go run ./ch07-modules-tooling/01-modules
//
// This whole repository is ONE module. Its go.mod file, at the top, says:
//
//	module github.com/delhiiitian/one-ninety-nine-go
//	go 1.23
//
// "module" is the name (import path) of the project.
// "go" is the minimum Go version the code needs.
// To start your own project:  mkdir myapp && cd myapp && go mod init example.com/myapp
package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("no build info available")
		return
	}
	fmt.Println("module path: ", info.Main.Path)
	fmt.Println("package path:", info.Path)
	fmt.Println("built with:  ", runtime.Version())
	fmt.Println("dependencies:", len(info.Deps), "(this module uses only the standard library)")
}
