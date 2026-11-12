package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var version = "dev"

func newHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"status": "healthy",
		})
	})

	mux.HandleFunc("GET /ready", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"status": "ready",
		})
	})

	mux.HandleFunc("GET /api/version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{
			"version": version,
		})
	})

	return mux
}

func writeJSON(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("response encoding failed: %v", err)
	}
}

func main() {
	if v := os.Getenv("APP_VERSION"); v != "" {
		version = v
	}

	server := &http.Server{
		Addr:              ":8080",
		Handler:           newHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("ReleaseFlow listening on port 8080")
	log.Fatal(server.ListenAndServe())
}
