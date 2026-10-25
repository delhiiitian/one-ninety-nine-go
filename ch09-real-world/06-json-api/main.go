// Chapter 9.6: A JSON REST API
// Run (demo, exits by itself): go run ./ch09-real-world/06-json-api
// Run (keep serving):          go run ./ch09-real-world/06-json-api -serve
//
//	curl localhost:8080/products
//	curl -X POST localhost:8080/products -d '{"name":"Mango","price":120}'
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Store is safe to use from many requests at once (each request runs in its own goroutine).
type Store struct {
	mu       sync.Mutex
	products map[int]Product
	nextID   int
}

func NewStore() *Store {
	s := &Store{products: map[int]Product{}, nextID: 1}
	s.Add(Product{Name: "Apple", Price: 30})
	s.Add(Product{Name: "Banana", Price: 10})
	return s
}

func (s *Store) Add(p Product) Product {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.ID = s.nextID
	s.nextID++
	s.products[p.ID] = p
	return p
}

func (s *Store) Get(id int) (Product, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.products[id]
	return p, ok
}

func (s *Store) All() []Product {
	s.mu.Lock()
	defer s.mu.Unlock()
	list := make([]Product, 0, len(s.products))
	for id := 1; id < s.nextID; id++ {
		if p, ok := s.products[id]; ok {
			list = append(list, p)
		}
	}
	return list
}

// writeJSON sends any value as JSON with a status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func routes(s *Store) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, s.All())
	})

	mux.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id")) // {id} from the pattern
		if err != nil {
			writeError(w, http.StatusBadRequest, "id must be a number")
			return
		}
		p, ok := s.Get(id)
		if !ok {
			writeError(w, http.StatusNotFound, "product not found")
			return
		}
		writeJSON(w, http.StatusOK, p)
	})

	mux.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {
		var p Product
		dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)) // max 1 MB
		dec.DisallowUnknownFields()
		if err := dec.Decode(&p); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
			return
		}
		if p.Name == "" || p.Price <= 0 {
			writeError(w, http.StatusUnprocessableEntity, "name and a positive price are required")
			return
		}
		writeJSON(w, http.StatusCreated, s.Add(p))
	})
	return mux
}

func main() {
	serve := flag.Bool("serve", false, "keep the server running on :8080")
	flag.Parse()
	handler := routes(NewStore())

	if *serve {
		log.Println("API on http://localhost:8080/products")
		log.Fatal(http.ListenAndServe(":8080", handler))
	}

	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	srv := &http.Server{Handler: handler}
	go srv.Serve(ln)
	defer srv.Close()
	base := "http://" + ln.Addr().String()

	call := func(method, path, body string) {
		req, _ := http.NewRequest(method, base+path, bytes.NewBufferString(body))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		out, _ := io.ReadAll(resp.Body)
		if body != "" {
			path += " " + body
		}
		fmt.Printf("%-6s %-14s -> %d %s", method, path, resp.StatusCode, out)
	}
	call("GET", "/products", "")
	call("POST", "/products", `{"name":"Mango","price":120}`)
	call("GET", "/products/3", "")
	call("GET", "/products/99", "")
	call("GET", "/products/abc", "")
	call("POST", "/products", `{"name":"","price":-1}`)
	call("POST", "/products", `{"name":"Kiwi","colour":"green"}`)
	call("DELETE", "/products/1", "")
}
