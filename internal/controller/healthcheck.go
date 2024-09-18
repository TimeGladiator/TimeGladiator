package controller

import (
	"fmt"
	"net/http"

	"github.com/TimeGladiator/TimeGladiator/internal/config"
)

// HealthcheckHandler handles request to /v1/healthcheck
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", cfg.Env)
	fmt.Fprintf(w, "version: %s\n", config.Version)
	// TODO: Check database connection
}
