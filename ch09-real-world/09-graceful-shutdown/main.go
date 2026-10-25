// Chapter 9.9: Graceful shutdown
// Run (demo, shuts itself down): go run ./ch09-real-world/09-graceful-shutdown
// Run (real server):             go run ./ch09-real-world/09-graceful-shutdown -serve
//
//	then press Ctrl+C while a request to http://localhost:8080/report is running.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serve := flag.Bool("serve", false, "run on :8080 until Ctrl+C")
	flag.Parse()
	log.SetFlags(0)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /report", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(300 * time.Millisecond) // a slow request that must not be cut off
		fmt.Fprintln(w, "monthly report ready")
	})
	srv := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}

	addr := "127.0.0.1:0"
	if *serve {
		addr = ":8080"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Run the server in the background.
	serverErr := make(chan error, 1)
	go func() {
		log.Println("serving on", ln.Addr())
		serverErr <- srv.Serve(ln)
	}()

	// 2. Wait for Ctrl+C or a stop signal from the system (Docker, Kubernetes send SIGTERM).
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if !*serve {
		go demoTraffic("http://"+ln.Addr().String(), stop)
	}

	select {
	case err := <-serverErr:
		log.Fatal("server failed:", err)
	case <-ctx.Done():
		log.Println("shutdown signal received: no new connections, finishing current requests...")
	}

	// 3. Give in-flight requests up to 5 seconds to finish.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Println("forced shutdown:", err)
	}
	if err := <-serverErr; !errors.Is(err, http.ErrServerClosed) {
		log.Println("unexpected:", err)
	}
	time.Sleep(20 * time.Millisecond) // let the demo client print its last line
	log.Println("clean exit: close database connections and flush logs here")
}

// demoTraffic starts a slow request, then "presses Ctrl+C" while it is still running.
func demoTraffic(base string, pressCtrlC func()) {
	go func() {
		resp, err := http.Get(base + "/report")
		if err != nil {
			log.Println("client error:", err)
			return
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("client got %d: %s", resp.StatusCode, body[:len(body)-1])
	}()
	time.Sleep(100 * time.Millisecond)
	log.Println("(demo) pressing Ctrl+C while /report is still running")
	pressCtrlC()
}
