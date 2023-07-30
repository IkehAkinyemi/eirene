package api

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *Server) setupRoutes() http.Handler {
	mux := pat.New()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		fileserver.ServeHTTP(w, r)
	})))

	mux.Get("/", http.HandlerFunc(s.redir2Home))
	mux.Get("/home", http.HandlerFunc(s.home))
	mux.Get("/posts", http.HandlerFunc(s.posts))
	mux.Get("/post/:id", http.HandlerFunc(s.post))

	return mux
}
