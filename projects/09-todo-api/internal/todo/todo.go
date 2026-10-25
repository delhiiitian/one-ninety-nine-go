// Package todo holds the to-do list rules and storage. It knows nothing about HTTP.
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid task")
)

// Task is one item on the list.
type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
}

// Patch lists the fields a client may change. nil means "leave as it is".
type Patch struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func cleanTitle(t string) (string, error) {
	t = strings.TrimSpace(t)
	switch {
	case t == "":
		return "", fmt.Errorf("%w: title is required", ErrInvalid)
	case len([]rune(t)) > 200:
		return "", fmt.Errorf("%w: title must be 200 characters or fewer", ErrInvalid)
	}
	return t, nil
}

// Store keeps tasks in memory and saves them to a JSON file after every change.
// A mutex makes it safe for many HTTP requests at once.
type Store struct {
	mu     sync.Mutex
	path   string // empty: memory only (used by tests)
	tasks  []Task
	nextID int
	now    func() time.Time
}

// Open loads the store from path, or starts empty if the file does not exist yet.
func Open(path string) (*Store, error) {
	s := &Store{path: path, nextID: 1, now: time.Now}
	if path == "" {
		return s, nil
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &s.tasks); err != nil {
		return nil, fmt.Errorf("%s is not valid JSON: %w", path, err)
	}
	for _, t := range s.tasks {
		s.nextID = max(s.nextID, t.ID+1)
	}
	return s, nil
}

// save must be called with the lock held. It writes a temp file and renames it (crash-safe).
func (s *Store) save() error {
	if s.path == "" {
		return nil
	}
	data, err := json.MarshalIndent(s.tasks, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}

// List returns tasks, optionally filtered by done state (nil = all).
func (s *Store) List(done *bool) []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := []Task{} // never nil, so it encodes as [] not null
	for _, t := range s.tasks {
		if done == nil || t.Done == *done {
			out = append(out, t)
		}
	}
	return out
}

func (s *Store) Get(id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i := s.index(id)
	if i < 0 {
		return Task{}, ErrNotFound
	}
	return s.tasks[i], nil
}

func (s *Store) Add(title string) (Task, error) {
	title, err := cleanTitle(title)
	if err != nil {
		return Task{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	t := Task{ID: s.nextID, Title: title, Created: s.now().UTC().Truncate(time.Second)}
	s.nextID++
	s.tasks = append(s.tasks, t)
	return t, s.save()
}

func (s *Store) Update(id int, p Patch) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i := s.index(id)
	if i < 0 {
		return Task{}, ErrNotFound
	}
	t := s.tasks[i]
	if p.Title != nil {
		title, err := cleanTitle(*p.Title)
		if err != nil {
			return Task{}, err
		}
		t.Title = title
	}
	if p.Done != nil {
		t.Done = *p.Done
	}
	s.tasks[i] = t
	return t, s.save()
}

func (s *Store) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	i := s.index(id)
	if i < 0 {
		return ErrNotFound
	}
	s.tasks = slices.Delete(s.tasks, i, i+1)
	return s.save()
}

func (s *Store) index(id int) int {
	return slices.IndexFunc(s.tasks, func(t Task) bool { return t.ID == id })
}
