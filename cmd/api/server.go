package api

import "net/http"

type Server struct {}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.home)
	return mux
}