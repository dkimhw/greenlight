package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// define config struct - to hold hold cofiguration settings
type config struct {
	port int
	env  string
}

// define app struct to hold dependencies for HTTP handlers, helpers, and middlewares.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port") // default val: 4000
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Writes log entries to the standard out stream
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare instance of app struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// declare HTTP server - uses servemux as the handler
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
