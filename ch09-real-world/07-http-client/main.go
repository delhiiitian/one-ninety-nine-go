// Chapter 9.7: Calling other services (HTTP client)
// Run: go run ./ch09-real-world/07-http-client
// A local "weather service" is started inside the program so it runs offline.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type Weather struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	Sky   string  `json:"sky"`
}

// fakeWeatherService: Pune works, "Slowtown" is slow, and the first call to Mumbai fails.
func fakeWeatherService() string {
	var mumbaiCalls atomic.Int32
	mux := http.NewServeMux()
	mux.HandleFunc("GET /weather/{city}", func(w http.ResponseWriter, r *http.Request) {
		city := r.PathValue("city")
		switch city {
		case "Slowtown":
			time.Sleep(2 * time.Second)
		case "Mumbai":
			if mumbaiCalls.Add(1) == 1 {
				http.Error(w, "temporarily overloaded", http.StatusServiceUnavailable)
				return
			}
		case "Atlantis":
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(Weather{city, 29.5, "partly cloudy"})
	})
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go http.Serve(ln, mux)
	return "http://" + ln.Addr().String()
}

// WeatherClient wraps http.Client with a base URL, a timeout and retries.
type WeatherClient struct {
	base string
	http *http.Client
}

var ErrNotFound = errors.New("city not found")

func (c *WeatherClient) Get(ctx context.Context, city string) (Weather, error) {
	var lastErr error
	for attempt := 1; attempt <= 3; attempt++ {
		w, err := c.getOnce(ctx, city)
		if err == nil || errors.Is(err, ErrNotFound) {
			return w, err // success, or an error that retrying will not fix
		}
		lastErr = err
		if attempt == 3 {
			break
		}
		wait := time.Duration(attempt*attempt) * 50 * time.Millisecond // back off: 50ms, 200ms
		fmt.Printf("    attempt %d failed (%v), retrying in %v\n", attempt, err, wait)
		select {
		case <-time.After(wait):
		case <-ctx.Done():
			return Weather{}, ctx.Err()
		}
	}
	return Weather{}, fmt.Errorf("gave up after 3 attempts: %w", lastErr)
}

func (c *WeatherClient) getOnce(ctx context.Context, city string) (Weather, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+"/weather/"+city, nil)
	if err != nil {
		return Weather{}, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		var ne net.Error
		if errors.As(err, &ne) && ne.Timeout() {
			return Weather{}, fmt.Errorf("timed out after %v", c.http.Timeout)
		}
		return Weather{}, err
	}
	defer resp.Body.Close() // always close the body

	switch {
	case resp.StatusCode == http.StatusNotFound:
		return Weather{}, ErrNotFound
	case resp.StatusCode != http.StatusOK:
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
		return Weather{}, fmt.Errorf("status %d: %s", resp.StatusCode, bytesTrim(msg))
	}

	var w Weather
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return Weather{}, fmt.Errorf("bad JSON from server: %w", err)
	}
	return w, nil
}

func bytesTrim(b []byte) string {
	for len(b) > 0 && (b[len(b)-1] == '\n' || b[len(b)-1] == ' ') {
		b = b[:len(b)-1]
	}
	return string(b)
}

func main() {
	client := &WeatherClient{
		base: fakeWeatherService(),
		http: &http.Client{Timeout: 500 * time.Millisecond}, // never use a client without a timeout
	}
	for _, city := range []string{"Pune", "Mumbai", "Atlantis", "Slowtown"} {
		fmt.Println("asking for", city)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		w, err := client.Get(ctx, city)
		cancel()
		if err != nil {
			fmt.Println("  error:", err)
			continue
		}
		fmt.Printf("  %s: %.1f C, %s\n", w.City, w.TempC, w.Sky)
	}
}
