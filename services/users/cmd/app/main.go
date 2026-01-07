package main

import (
	"log"
	"os"
	"users/internal/app"
)

func main() {
	dsn := os.Getenv("USERS_DB_DSN")
	if dsn == "" {
		log.Fatal("USERS_DB_DSN is not set")
	}
	application, err := app.New(dsn)
	if err != nil {
		log.Fatalf("application.New: %s", err)
	}
	if err := application.Run(); err != nil {
		log.Fatalf("application.Run: %s", err)
	}
}
