package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Cnf struct {
	Port string
	Env  string
	DatabaseUrl string
}

func MustLoad() Cnf {
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		panic("PORT is required")
	}

	env := os.Getenv("ENV")

	if env == "" {
		panic("ENV is required")
	}
    
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		panic("DATABASE_URL is required")
	}

	return Cnf{
		Port: port,
		Env:  env,
		DatabaseUrl: databaseUrl,
	}

}
