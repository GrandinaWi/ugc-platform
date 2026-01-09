package routes

import (
	"net/http"
	handler2 "posts/internal/routes/handler"
	"posts/internal/service"
)

func NewRouter(postsService service.Service) http.Handler {
	mux := http.NewServeMux()
	handler := handler2.NewPostsHandler(postsService)
	mux.HandleFunc("POST /posts", handler.CreatePost)
	mux.HandleFunc("GET /posts", handler.GetPosts)
	mux.HandleFunc("GET /posts/{post_id}", handler.GetPost)
	mux.HandleFunc("DELETE /posts/{post_id}", handler.DeletePost)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}
