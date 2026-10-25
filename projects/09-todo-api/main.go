// Project 9: To-do List REST API
//
//	go run ./projects/09-todo-api                         # listens on :8080, saves to tasks.json
//	go run ./projects/09-todo-api -addr :9000 -data /tmp/tasks.json
//	TODO_ADDR=:9000 go run ./projects/09-todo-api        # environment variables work too
//	go test ./projects/09-todo-api/...
//
// Try it from another terminal:
//
//	curl -s localhost:8080/tasks
//	curl -s -X POST localhost:8080/tasks -d '{"title":"Buy milk"}'
//	curl -s -X PATCH localhost:8080/tasks/1 -d '{"done":true}'
//	curl -s -X DELETE localhost:8080/tasks/1
package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/delhiiitian/one-ninety-nine-go/projects/09-todo-api/internal/api"
	"github.com/delhiiitian/one-ninety-nine-go/projects/09-todo-api/internal/todo"
)

func main() {
	addr := flag.String("addr", env("TODO_ADDR", ":8080"), "address to listen on")
	data := flag.String("data", env("TODO_DATA", "tasks.json"), "file to save tasks in")
	jsonLogs := flag.Bool("json-logs", false, "write logs as JSON")
	flag.Parse()

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, nil)
	if *jsonLogs {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	}
	log := slog.New(handler)

	store, err := todo.Open(*data)
	if err != nil {
		log.Error("cannot open data file", "err", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:              *addr,
		Handler:           api.Handler(store, log),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Info("to-do API listening", "addr", *addr, "data", *data)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server failed", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	log.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Error("forced shutdown", "err", err)
	}
	log.Info("bye")
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
