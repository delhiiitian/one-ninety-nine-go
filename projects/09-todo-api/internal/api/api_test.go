package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/delhiiitian/one-ninety-nine-go/projects/09-todo-api/internal/todo"
)

func newServer(t *testing.T, path string) *httptest.Server {
	t.Helper()
	store, err := todo.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(Handler(store, slog.New(slog.NewTextHandler(io.Discard, nil))))
	t.Cleanup(srv.Close)
	return srv
}

func do(t *testing.T, method, url, body string) (int, string) {
	t.Helper()
	req, _ := http.NewRequest(method, url, strings.NewReader(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, strings.TrimSpace(string(b))
}

func TestTaskLifecycle(t *testing.T) {
	srv := newServer(t, "")
	u := srv.URL

	steps := []struct {
		method, path, body string
		status             int
		contains           string
	}{
		{"GET", "/tasks", "", 200, "[]"},
		{"POST", "/tasks", `{"title":"  Buy milk  "}`, 201, `"title":"Buy milk"`},
		{"POST", "/tasks", `{"title":"Pay rent"}`, 201, `"id":2`},
		{"POST", "/tasks", `{"title":""}`, 422, "title is required"},
		{"POST", "/tasks", `{"name":"x"}`, 400, "unknown field"},
		{"PATCH", "/tasks/1", `{"done":true}`, 200, `"done":true`},
		{"GET", "/tasks?done=true", "", 200, "Buy milk"},
		{"GET", "/tasks?done=maybe", "", 400, "true or false"},
		{"GET", "/tasks/9", "", 404, "not found"},
		{"GET", "/tasks/abc", "", 400, "positive number"},
		{"DELETE", "/tasks/2", "", 204, ""},
		{"DELETE", "/tasks/2", "", 404, "not found"},
		{"PUT", "/tasks/1", "{}", 405, ""},
	}
	for _, s := range steps {
		status, body := do(t, s.method, u+s.path, s.body)
		if status != s.status || !strings.Contains(body, s.contains) {
			t.Errorf("%s %s %s: got %d %s, want %d containing %q",
				s.method, s.path, s.body, status, body, s.status, s.contains)
		}
	}
}

func TestTasksSurviveRestart(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	srv := newServer(t, path)
	do(t, "POST", srv.URL+"/tasks", `{"title":"Water plants"}`)
	srv.Close()

	srv2 := newServer(t, path) // a "restarted" server reading the same file
	_, body := do(t, "GET", srv2.URL+"/tasks", "")
	var tasks []todo.Task
	if err := json.Unmarshal([]byte(body), &tasks); err != nil || len(tasks) != 1 || tasks[0].Title != "Water plants" {
		t.Fatalf("after restart got %s (err %v)", body, err)
	}
	_, body = do(t, "POST", srv2.URL+"/tasks", `{"title":"Next"}`)
	if !strings.Contains(body, `"id":2`) {
		t.Fatalf("IDs must continue after restart, got %s", body)
	}
}
