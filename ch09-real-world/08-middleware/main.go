// Chapter 9.8: Middleware (wrapping handlers with shared behaviour)
// Run: go run ./ch09-real-world/08-middleware
package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

// A middleware takes a handler and returns a new handler that wraps it.
type Middleware func(http.Handler) http.Handler

// statusRecorder remembers the status code a handler wrote.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// Logging records every request: method, path, status and time taken.
func Logging(log *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(rec, r)
			log.Info("request", "method", r.Method, "path", r.URL.Path,
				"status", rec.status, "slow", time.Since(start) > 100*time.Millisecond)
		})
	}
}

// Recover turns a panic in any handler into a 500 response instead of a crash.
func Recover(log *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("panic", "err", err, "path", r.URL.Path)
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// RequireKey blocks requests that do not carry the right API key.
func RequireKey(key string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-API-Key") != key {
				http.Error(w, "missing or wrong API key", http.StatusUnauthorized)
				return // next is never called
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Chain applies middlewares so the first one listed is the outermost.
func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(g []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /public", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "anyone can see this")
	})
	mux.Handle("GET /admin", RequireKey("s3cret")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "sales today: Rs 48,250")
	})))
	mux.HandleFunc("GET /buggy", func(w http.ResponseWriter, r *http.Request) {
		var m map[string]int
		m["boom"] = 1 // panics: nil map
	})

	handler := Chain(mux, Logging(log), Recover(log))

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	srv := &http.Server{Handler: handler}
	go srv.Serve(ln)
	defer srv.Close()
	base := "http://" + ln.Addr().String()

	for _, c := range []struct{ path, key string }{
		{"/public", ""}, {"/admin", ""}, {"/admin", "s3cret"}, {"/buggy", ""}, {"/missing", ""},
	} {
		req, _ := http.NewRequest("GET", base+c.path, nil)
		if c.key != "" {
			req.Header.Set("X-API-Key", c.key)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("  -> client got %d: %s\n", resp.StatusCode, trim(body))
	}
}

func trim(b []byte) string {
	s := string(b)
	for len(s) > 0 && s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return s
}
