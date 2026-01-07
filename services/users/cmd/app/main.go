package main

import (
	"log"
	"os"
	"users/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
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
