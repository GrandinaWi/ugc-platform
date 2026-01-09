package routes

import (
	"net/http"
	handler2 "users/internal/routes/handler"
	service "users/internal/service"
)

func NewRouter(userService service.Service) http.Handler {
	mux := http.NewServeMux()
	handler := handler2.NewUsersHandler(userService)
	mux.HandleFunc("POST /register", handler.UserRegisterHandler)
	mux.HandleFunc("POST /login", handler.UserLoginHandler)
	mux.HandleFunc("GET /user", handler.UserInfoHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	return mux
}
