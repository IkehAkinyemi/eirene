package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) setupRoutes() http.Handler {
	mux := httprouter.New()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileserver := http.FileServer(http.Dir("./ui/static/"))
	mux.Handler(http.MethodGet, "/static/", http.StripPrefix("/static", fileserver))

	return mux
}