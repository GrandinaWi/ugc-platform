package app

import (
	"comment/internal/repository/postgres"
	"comment/internal/routes"
	postgres2 "comment/internal/service/postgres"
	"context"
	"database/sql"
	"net/http"
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

	commentsRepo := postgres.NewRepository(db)
	commentsService := postgres2.NewService(commentsRepo)

	router := routes.NewRouter(commentsService)

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
