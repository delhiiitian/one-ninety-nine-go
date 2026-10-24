package checker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// testServer gives us a real HTTP server inside the test, with no internet needed.
func testServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
	mux.HandleFunc("/missing", http.NotFound)
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(time.Second):
		case <-r.Context().Done():
		}
	})
	return httptest.NewServer(mux)
}

func TestCheckAll(t *testing.T) {
	srv := testServer()
	defer srv.Close()

	c := New(3, 200*time.Millisecond)
	urls := []string{srv.URL + "/ok", srv.URL + "/missing", srv.URL + "/slow", "http://bad host"}
	got := c.CheckAll(context.Background(), urls)

	if len(got) != 4 {
		t.Fatalf("got %d results", len(got))
	}
	for i, u := range urls {
		if got[i].URL != u {
			t.Errorf("result %d is for %s, want %s (order must be kept)", i, got[i].URL, u)
		}
	}
	if !got[0].Up() || got[0].Status != 200 {
		t.Errorf("/ok: %+v", got[0])
	}
	if got[1].Up() || got[1].Status != 404 {
		t.Errorf("/missing: %+v", got[1])
	}
	if got[2].Err == nil {
		t.Errorf("/slow should time out: %+v", got[2])
	}
	if got[3].Err == nil {
		t.Errorf("bad URL should fail: %+v", got[3])
	}
}

func TestWorkersRunInParallel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	}))
	defer srv.Close()

	urls := make([]string, 8)
	for i := range urls {
		urls[i] = srv.URL
	}
	start := time.Now()
	New(8, time.Second).CheckAll(context.Background(), urls)
	if d := time.Since(start); d > 500*time.Millisecond {
		t.Fatalf("8 checks with 8 workers took %v; they are not running in parallel", d)
	}
}
