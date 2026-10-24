package main

import (
	"net"
	"net/http"
	"time"
)

// startDemoServer runs a small local website with fast, slow, broken and
// missing pages, so the checker can be tried without the internet.
// It returns the base URL, for example http://127.0.0.1:53211.
func startDemoServer() (string, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) { // "{$}": only the exact path /
		w.Write([]byte("home"))
	})
	mux.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(120 * time.Millisecond)
		w.Write([]byte("shop"))
	})
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // too slow: will time out
	})
	mux.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "database down", http.StatusServiceUnavailable)
	})
	mux.HandleFunc("/old-page", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0") // port 0: pick any free port
	if err != nil {
		return "", err
	}
	go http.Serve(ln, mux)
	return "http://" + ln.Addr().String(), nil
}
