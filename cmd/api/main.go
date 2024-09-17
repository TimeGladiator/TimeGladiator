package main

import (
	"github.com/TimeGladiator/TimeGladiator/internal/config"
	"github.com/TimeGladiator/TimeGladiator/internal/route"
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
