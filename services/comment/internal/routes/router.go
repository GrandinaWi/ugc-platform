package routes

import (
	handler2 "comment/internal/routes/handler"
	"comment/internal/service"
	"net/http"
)

func NewRouter(commentsService service.Service) http.Handler {
	mux := http.NewServeMux()
	handler := handler2.NewHandler(commentsService)
	mux.HandleFunc("POST /posts/{post_id}/comments", handler.CreateHandler)
	mux.HandleFunc("GET /posts/{post_id}/comments", handler.GetAllHandler)
	mux.HandleFunc("GET /comments/{id}", handler.GetByIDHandler)
	mux.HandleFunc("DELETE /comments/{id}", handler.DeleteHandler)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}
