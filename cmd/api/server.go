package main

import (
	"fmt"
	"github.com/buker/go-basic-template/internal/config"
	"net/http"
	"time"
)

// Serve serves an HTTP server
func Serve(cfg config.Config, handler http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP Server
	fmt.Printf("Starting %s server on :%d\n", cfg.Env, cfg.Port)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
