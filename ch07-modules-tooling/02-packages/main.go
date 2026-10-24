// Chapter 7.2: Packages and exported names
// Run: go run ./ch07-modules-tooling/02-packages
package main

import (
	"fmt"

	// Import path = module path + folder path
	"github.com/delhiiitian/one-ninety-nine-go/ch07-modules-tooling/02-packages/greeting"
)

func main() {
	fmt.Println(greeting.Hello("asha", "hi"))
	fmt.Println(greeting.Hello("  carlos ", "es"))
	fmt.Println(greeting.Hello("", "en"))
	fmt.Println("default name is", greeting.DefaultName)

	// greeting.word("hi") // compile error: word is not exported
}
