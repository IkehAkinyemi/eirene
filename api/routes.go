package api

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *Server) setupRoutes() http.Handler {
	mux := pat.New()

	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	mux.Get("/", http.HandlerFunc(s.redir2Home))
	mux.Get("/home", http.HandlerFunc(s.home))
	mux.Get("/posts", http.HandlerFunc(s.posts))
	mux.Get("/post/:id", http.HandlerFunc(s.post))

	return mux
}
