package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mehedishishir0/levox-api/internal/config"
	"github.com/mehedishishir0/levox-api/internal/db"
	"github.com/mehedishishir0/levox-api/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	_, err := db.Connect(cfg.DatabaseUrl)

	mux := http.NewServeMux()

	if err != nil {
		log.Fatalf("main.Connect error: %v", err)
	}

	fmt.Printf("Connected to database: %s\n", cfg.DatabaseUrl)

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
