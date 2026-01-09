package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"posts/internal/model"
	"posts/internal/service"
	"strconv"
)

type PostsHandler struct {
	service service.Service
}

func NewUsersHandler(service service.Service) *PostsHandler {
	return &PostsHandler{service: service}
}

func (s *PostsHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := s.service.PostsGet(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (s *PostsHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("post_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := s.service.PostGet(ctx, id)
	if errors.Is(err, model.PostNotExist) {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(post)
}
func (s *PostsHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req struct {
		AuthorID int64  `json:"author_id"`
		Title    string `json:"title"`
		Content  string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	if req.AuthorID == 0 || req.Title == "" || req.Content == "" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	id, err := s.service.PostCreate(ctx, req.AuthorID, req.Title, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 for created
	_ = json.NewEncoder(w).Encode(id)

}
func (s *PostsHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("post_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.service.PostDelete(ctx, id)
	if err != nil {
		if errors.Is(err, model.PostNotExist) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
