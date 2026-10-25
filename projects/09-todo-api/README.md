# Project 9: To-do List REST API

A small but production-shaped web service: a JSON REST API for a to-do list, with
validation, proper status codes, structured request logs, panic recovery, tasks saved to
a file so they survive restarts, configuration by flags or environment variables, and
graceful shutdown. Any web or mobile front end could use it.

Book: *Go in 199 Pages*, Chapter 9 project.

## Run it

```bash
go run ./projects/09-todo-api -data /tmp/tasks.json
```

In another terminal:

```bash
curl -s localhost:8080/tasks
curl -s -X POST localhost:8080/tasks -d '{"title":"Buy milk"}'
curl -s -X PATCH localhost:8080/tasks/1 -d '{"done":true}'
curl -s "localhost:8080/tasks?done=true"
curl -s -X DELETE localhost:8080/tasks/1 -i
```

Press Ctrl+C to stop; running requests finish first.

## Endpoints

| Method | Path | Body | Success |
| --- | --- | --- | --- |
| GET | `/healthz` | | 200 `{"status":"ok"}` |
| GET | `/tasks?done=true` | | 200 list (filter optional) |
| POST | `/tasks` | `{"title":"..."}` | 201 task, `Location` header |
| GET | `/tasks/{id}` | | 200 task |
| PATCH | `/tasks/{id}` | `{"title":"..."}` and/or `{"done":true}` | 200 task |
| DELETE | `/tasks/{id}` | | 204 |

Errors are JSON `{"error":"..."}` with 400 (bad input), 404 (no such task), 405 (wrong
method), 422 (fails validation) or 500.

## Layout

| Path | Job |
| --- | --- |
| `main.go` | Configuration, logger, server timeouts, graceful shutdown |
| `internal/todo/todo.go` | The `Task` model, validation, and a thread-safe store saved to JSON |
| `internal/api/handlers.go` | Routes, request decoding, error-to-status mapping |
| `internal/api/middleware.go` | Request logging and panic recovery |
| `internal/api/api_test.go` | End-to-end tests with `httptest`, including a restart test |

## Ideas to extend

- Add due dates and a `GET /tasks?due=today` filter.
- Add an API key middleware (see lesson 9.8).
- Replace the JSON file with SQLite by writing a second store with the same methods.
