package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mehedishishir0/levox-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"message":"all ok"}`))
	})

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
