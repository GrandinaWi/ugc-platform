package handler

import (
	"comment/internal/auth"
	"comment/internal/model"
	"comment/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	service service.Service
}

func NewHandler(service service.Service) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var commentID int64
	postIDStr := r.PathValue("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil || postID <= 0 {
		http.Error(w, "invalid post_id", http.StatusBadRequest)
		return
	}
	var req struct {
		ParentID *int64 `json:"parent_id"`
		Content  string `json:"content"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	comment := model.Comment{
		PostID:   postID,
		UserID:   userID,
		ParentID: req.ParentID,
		Content:  req.Content,
	}
	commentID, err = h.service.Create(ctx, &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id": commentID,
	})
}
func (h *CommentHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postIDStr := r.PathValue("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil || postID <= 0 {
		http.Error(w, "invalid post_id", http.StatusBadRequest)
		return
	}

	comments, err := h.service.GetByAll(ctx, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}
func (h *CommentHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	userID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.service.DeleteByID(ctx, id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
func (h *CommentHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var comment *model.Comment
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	comment, err = h.service.GetByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)

}
