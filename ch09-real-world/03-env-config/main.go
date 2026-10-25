// Chapter 9.3: Environment variables and configuration
// Run: go run ./ch09-real-world/03-env-config
// Try: APP_PORT=9090 APP_DEBUG=true go run ./ch09-real-world/03-env-config
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Config holds every setting the app needs, in one place.
type Config struct {
	Port        int           `json:"port"`
	DatabaseURL string        `json:"database_url"`
	Debug       bool          `json:"debug"`
	Timeout     time.Duration `json:"-"`
	TimeoutText string        `json:"timeout"`
}

// Load builds the config in layers: defaults, then a JSON file, then environment variables.
// Later layers win. This is how most production services are configured.
func Load(file string) (Config, error) {
	cfg := Config{Port: 8080, DatabaseURL: "postgres://localhost/shop", TimeoutText: "5s"}

	if data, err := os.ReadFile(file); err == nil {
		if err := json.Unmarshal(data, &cfg); err != nil {
			return cfg, fmt.Errorf("config file %s: %w", file, err)
		}
	}

	if v, ok := os.LookupEnv("APP_PORT"); ok {
		port, err := strconv.Atoi(v)
		if err != nil {
			return cfg, fmt.Errorf("APP_PORT=%q is not a number", v)
		}
		cfg.Port = port
	}
	if v := os.Getenv("APP_DATABASE_URL"); v != "" {
		cfg.DatabaseURL = v
	}
	if v := os.Getenv("APP_DEBUG"); v != "" {
		cfg.Debug, _ = strconv.ParseBool(v)
	}
	if v := os.Getenv("APP_TIMEOUT"); v != "" {
		cfg.TimeoutText = v
	}

	d, err := time.ParseDuration(cfg.TimeoutText)
	if err != nil {
		return cfg, fmt.Errorf("timeout %q: %w", cfg.TimeoutText, err)
	}
	cfg.Timeout = d
	return cfg, nil
}

func main() {
	// A config file for the demo
	file := filepath.Join(os.TempDir(), "shop-config.json")
	os.WriteFile(file, []byte(`{"port": 3000, "timeout": "10s"}`), 0o644)
	defer os.Remove(file)

	cfg, err := Load(file)
	if err != nil {
		fmt.Println("config error:", err)
		os.Exit(1)
	}
	fmt.Printf("port=%d debug=%v timeout=%v db=%s\n", cfg.Port, cfg.Debug, cfg.Timeout, cfg.DatabaseURL)

	// Setting a variable for this process (and the programs it starts)
	os.Setenv("APP_PORT", "4000")
	cfg, _ = Load(file)
	fmt.Println("after APP_PORT=4000, port is", cfg.Port)

	// Secrets belong in the environment, never in code or in Git.
	if os.Getenv("PAYMENT_API_KEY") == "" {
		fmt.Println("PAYMENT_API_KEY is not set (that's fine for this demo)")
	}
}
