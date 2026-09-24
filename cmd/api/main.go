package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mehedishishir0/levox-api/internal/config"
	"github.com/mehedishishir0/levox-api/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Healthz)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("server is starting on port http://localhost:%s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server not running on :8080 err: %v", err)
	}
}
