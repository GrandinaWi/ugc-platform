package router

import (
	"api-gateway/internal/auth"
	"api-gateway/internal/proxy"
	"net/http"
)

func New(secret []byte, usersURL string, postsURL string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/register", proxy.New(usersURL))
	mux.Handle("/login", proxy.New(usersURL))
	mux.Handle("/user", auth.Middleware(secret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		proxy.New(usersURL).ServeHTTP(w, r)
	})))

	mux.Handle("/posts", proxy.New(postsURL))
	mux.Handle("/posts/", proxy.New(postsURL))

	return mux

}
