package main

import (
	"log"
	"os"
	"users/internal/app"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	application, err := app.New(dsn)
	if err != nil {
		log.Fatalf("application.New: %s", err)
	}
	if err := application.Run(); err != nil {
		log.Fatalf("application.Run: %s", err)
	}
}
