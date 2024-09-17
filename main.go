package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type config struct {
	port int
	env  string
}

var cfg config

// Read value of command-line flags
func init() {
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/hello", helloHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP Server
	fmt.Printf("Starting %s server on :%d\n", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from main.go!")
}
