package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mehedishishir0/levox-api/internal/config"
)

func main() {

	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: up or down")
		os.Exit(1)
	}

	cfg := config.MustLoad()

	m, err := migrate.New("file://migrations", cfg.DatabaseUrl)
	if err != nil {
		fmt.Printf("Failed to create migrate instance: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("Migration up failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migration up completed successfully.")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			fmt.Printf("Migration down failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Migration down completed successfully.")
	default:
		m.Close()
		fmt.Println("Invalid command. Please use 'up' or 'down'")
		os.Exit(1)
	}
}
