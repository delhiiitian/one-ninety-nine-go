// Package api turns HTTP requests into calls on the todo store.
package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/delhiiitian/one-ninety-nine-go/projects/09-todo-api/internal/todo"
)

type API struct {
	store *todo.Store
	log   *slog.Logger
}

// Handler returns the full HTTP handler: routes wrapped in middleware.
func Handler(store *todo.Store, log *slog.Logger) http.Handler {
	a := &API{store: store, log: log}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /tasks", a.list)
	mux.HandleFunc("POST /tasks", a.create)
	mux.HandleFunc("GET /tasks/{id}", a.get)
	mux.HandleFunc("PATCH /tasks/{id}", a.update)
	mux.HandleFunc("DELETE /tasks/{id}", a.delete)
	return Chain(mux, RequestLog(log), Recover(log))
}

// GET /tasks?done=true|false
func (a *API) list(w http.ResponseWriter, r *http.Request) {
	var filter *bool
	if v := r.URL.Query().Get("done"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			writeError(w, http.StatusBadRequest, "done must be true or false")
			return
		}
		filter = &b
	}
	writeJSON(w, http.StatusOK, a.store.List(filter))
}

// POST /tasks  {"title": "..."}
func (a *API) create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	if !decode(w, r, &body) {
		return
	}
	t, err := a.store.Add(body.Title)
	if err != nil {
		a.fail(w, err)
		return
	}
	w.Header().Set("Location", "/tasks/"+strconv.Itoa(t.ID))
	writeJSON(w, http.StatusCreated, t)
}

// GET /tasks/{id}
func (a *API) get(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	t, err := a.store.Get(id)
	if err != nil {
		a.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// PATCH /tasks/{id}  {"done": true} or {"title": "..."}
func (a *API) update(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	var p todo.Patch
	if !decode(w, r, &p) {
		return
	}
	t, err := a.store.Update(id, p)
	if err != nil {
		a.fail(w, err)
		return
	}
	writeJSON(w, http.StatusOK, t)
}

// DELETE /tasks/{id}
func (a *API) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	if err := a.store.Delete(id); err != nil {
		a.fail(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// fail maps store errors to HTTP status codes in one place.
func (a *API) fail(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, todo.ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, todo.ErrInvalid):
		writeError(w, http.StatusUnprocessableEntity, err.Error())
	default:
		a.log.Error("store failure", "err", err)
		writeError(w, http.StatusInternalServerError, "internal error")
	}
}

func pathID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		writeError(w, http.StatusBadRequest, "id must be a positive number")
		return 0, false
	}
	return id, true
}

func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
