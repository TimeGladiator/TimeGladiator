package main

import (
	"github.com/buker/go-basic-template/internal/config"
	"github.com/buker/go-basic-template/internal/route"
	"log"
)

func main() {
	// Load config
	cfg := config.Load()

	// Register routes
	mux := route.RegisterRoutes()

	// Start HTTP server.
	err := Serve(cfg, mux)
	log.Fatal(err)
}
