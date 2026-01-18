package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	handler := router.New(
		cfg.JWTSecret,
		cfg.UsersAPI,
		cfg.PostsAPI,
		cfg.CommentsAPI,
	)
	log.Println("Gateway listening on:", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
