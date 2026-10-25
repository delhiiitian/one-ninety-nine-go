// Chapter 9.4: Structured logging with log/slog
// Run: go run ./ch09-real-world/04-logging
package main

import (
	"errors"
	"log"
	"log/slog"
	"os"
	"time"
)

func main() {
	// 1. The classic log package: simple lines with a timestamp.
	log.SetFlags(0) // no timestamp, so the book's output is stable
	log.SetPrefix("[shop] ")
	log.Println("server starting")

	// 2. slog: structured logs as key=value pairs that machines can search.
	text := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelDebug,
		ReplaceAttr: dropTime, // keep the demo output stable
	}))
	text.Info("order placed", "order_id", 1042, "amount", 1499.0, "items", 3)
	text.Debug("cache lookup", "key", "user:42", "hit", false)
	text.Warn("slow payment gateway", "latency", 2300*time.Millisecond)
	text.Error("payment failed", "order_id", 1043, "err", errors.New("card declined"))

	// 3. JSON logs: what log collectors (Datadog, Loki, CloudWatch) prefer.
	jsonLog := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: dropTime}))
	jsonLog.Info("user signed up", "user_id", 77, "plan", "pro")

	// 4. Loggers with context attached: every line from this request carries its ID.
	reqLog := text.With("request_id", "a1b2c3", "path", "/checkout")
	reqLog.Info("request started")
	reqLog.Info("request finished", "status", 200, "took", 42*time.Millisecond)

	// 5. Levels filter noise: at Info level, Debug lines are dropped.
	quiet := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, ReplaceAttr: dropTime}))
	quiet.Debug("you will not see this")
	quiet.Info("you will see this")
}

// dropTime removes the timestamp. Real programs keep it!
func dropTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}
