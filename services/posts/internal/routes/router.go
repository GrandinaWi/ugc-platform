package routes

import (
	"net/http"
	handler2 "posts/internal/routes/handler"
	"posts/internal/service"
)

func NewRouter(postsService service.Service) http.Handler {
	mux := http.NewServeMux()
	handler := handler2.NewPostsHandler(postsService)
	mux.HandleFunc("POST /create", handler.CreatePost)
	mux.HandleFunc("POST /delete", handler.DeletePost)
	mux.HandleFunc("GET /posts/", handler.GetPost)
	mux.HandleFunc("GET /posts", handler.GetPosts)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}
