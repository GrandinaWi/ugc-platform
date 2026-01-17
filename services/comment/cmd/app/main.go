package main

import (
	"comment/internal/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
