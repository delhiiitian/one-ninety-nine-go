// Chapter 9.5: An HTTP server
// Run (demo, exits by itself): go run ./ch09-real-world/05-http-server
// Run (keep serving):          go run ./ch09-real-world/05-http-server -serve
//
//	then open http://localhost:8080/ and http://localhost:8080/hello?name=Asha
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go 199 bakery!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // ?name=Asha
	if name == "" {
		name = "stranger"
	}
	fmt.Fprintf(w, "Hello, %s! You asked for %s with %s.\n", name, r.URL.Path, r.Method)
}

func menu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Menu</h1><ul><li>Croissant Rs 90</li><li>Chai Rs 30</li></ul>")
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home) // exactly "/"
	mux.HandleFunc("GET /hello", hello)
	mux.HandleFunc("GET /menu", menu)
	return mux
}

func main() {
	serve := flag.Bool("serve", false, "keep the server running on :8080")
	flag.Parse()

	srv := &http.Server{Handler: newMux(), ReadHeaderTimeout: 5 * time.Second}

	if *serve {
		srv.Addr = ":8080"
		log.Println("listening on http://localhost:8080 (Ctrl+C to stop)")
		log.Fatal(srv.ListenAndServe())
	}

	// Demo mode: listen on a free port, call ourselves, then stop.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go srv.Serve(ln)
	defer srv.Close()
	base := "http://" + ln.Addr().String()

	for _, path := range []string{"/", "/hello?name=Asha", "/menu", "/nope"} {
		resp, err := http.Get(base + path)
		if err != nil {
			log.Fatal(err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("GET %-17s -> %d %s", path, resp.StatusCode, body)
		if len(body) == 0 || body[len(body)-1] != '\n' {
			fmt.Println()
		}
	}
}
