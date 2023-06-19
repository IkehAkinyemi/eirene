package api

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (s *Server) setupRoutes() http.Handler {
	mux := pat.New()

	// mux.Get("/landing-page", http.HandlerFunc(landingPage))

	// Create a file server which serves files out of the "./ui/static" directory.
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileserver))

	mux.Get("/home", http.HandlerFunc(s.home))
	mux.Get("/posts", http.HandlerFunc(s.posts))
	mux.Get("/post", http.HandlerFunc(s.post))

	return mux
}
