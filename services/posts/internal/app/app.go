package app

import (
	"context"
	"database/sql"
	"net/http"
	"posts/internal/repository/postgres"
	"posts/internal/routes"
	postgres2 "posts/internal/service/postgres"
	"time"
)

type App struct {
	server *http.Server
}

func New(dsn string) (*App, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	postsRepo := postgres.NewRepository(db)
	postsService := postgres2.NewService(postsRepo)

	router := routes.NewRouter(postsService)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &App{server: server}, nil
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
