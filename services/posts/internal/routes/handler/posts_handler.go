package handler

import (
	"posts/internal/service"
)

type PostsHandler struct {
	service service.Service
}

func NewUsersHandler(service service.Service) *PostsHandler {
	return &PostsHandler{service: service}
}
