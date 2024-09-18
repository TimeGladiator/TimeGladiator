package main

import (
	"log"

	"github.com/TimeGladiator/TimeGladiator/internal/config"
	"github.com/TimeGladiator/TimeGladiator/internal/route"
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
