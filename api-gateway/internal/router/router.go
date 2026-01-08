package router

import (
	"api-gateway/internal/auth"
	"api-gateway/internal/proxy"
	"net/http"
)

func New(secret []byte, usersURL string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/register", proxy.New(usersURL))
	mux.Handle("/login", proxy.New(usersURL))
	mux.Handle("/user", auth.Middleware(secret, proxy.New(usersURL)))

	return mux

}
