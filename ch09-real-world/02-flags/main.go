// Chapter 9.2: Command-line flags
// Run:  go run ./ch09-real-world/02-flags
// Try:  go run ./ch09-real-world/02-flags -width 800 -format png -quality 70 -v photo1.jpg photo2.jpg
// Help: go run ./ch09-real-world/02-flags -h
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Each flag has a name, a default value and help text.
	width := flag.Int("width", 1024, "resize images to this width in pixels")
	format := flag.String("format", "jpg", "output format: jpg, png or webp")
	quality := flag.Int("quality", 85, "compression quality 1-100")
	verbose := flag.Bool("v", false, "print details for each file")
	timeout := flag.Duration("timeout", 30*time.Second, "stop after this long")

	// Flags can also fill an existing variable.
	var outDir string
	flag.StringVar(&outDir, "out", "resized", "folder for the results")

	flag.Parse() // reads os.Args; prints help and exits on -h

	// Validate like any other input.
	switch *format {
	case "jpg", "png", "webp":
	default:
		fmt.Fprintf(os.Stderr, "unknown -format %q\n", *format)
		flag.Usage()
		os.Exit(2)
	}
	if *quality < 1 || *quality > 100 {
		fmt.Fprintln(os.Stderr, "-quality must be 1-100")
		os.Exit(2)
	}

	files := flag.Args() // whatever is left after the flags
	if len(files) == 0 {
		files = []string{"example.jpg"}
	}

	fmt.Printf("settings: width=%d format=%s quality=%d out=%s timeout=%v\n",
		*width, *format, *quality, outDir, *timeout)
	for _, f := range files {
		name := strings.TrimSuffix(f, ".jpg") + "." + *format
		if *verbose {
			fmt.Printf("  %s -> %s/%s (%dpx, q%d)\n", f, outDir, name, *width, *quality)
		} else {
			fmt.Println("  would resize", f)
		}
	}
}
