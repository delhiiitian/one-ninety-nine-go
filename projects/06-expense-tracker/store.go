package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Store keeps all expenses in one JSON file.
type Store struct {
	Path     string    `json:"-"`
	NextID   int       `json:"next_id"`
	Expenses []Expense `json:"expenses"`
}

// Load reads the file. A missing file is not an error: it means "no expenses yet".
func Load(path string) (*Store, error) {
	s := &Store{Path: path, NextID: 1}
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, fmt.Errorf("load %s: %w", path, err)
	}
	if err := json.Unmarshal(data, s); err != nil {
		return nil, fmt.Errorf("load %s: file is not valid JSON: %w", path, err)
	}
	return s, nil
}

// Save writes to a temporary file first and then renames it over the real one.
// If the computer crashes halfway, the old file is still intact.
func (s *Store) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	if dir := filepath.Dir(s.Path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	tmp := s.Path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return fmt.Errorf("save: %w", err)
	}
	return os.Rename(tmp, s.Path)
}

func (s *Store) Add(e Expense) Expense {
	e.ID = s.NextID
	s.NextID++
	s.Expenses = append(s.Expenses, e)
	return e
}

func (s *Store) Delete(id int) error {
	for i, e := range s.Expenses {
		if e.ID == id {
			s.Expenses = append(s.Expenses[:i], s.Expenses[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no expense with ID %d", id)
}
