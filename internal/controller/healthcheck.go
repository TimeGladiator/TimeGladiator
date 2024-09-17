package controller

import (
	"fmt"
	"github.com/buker/go-basic-template/internal/config"
	"net/http"
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
