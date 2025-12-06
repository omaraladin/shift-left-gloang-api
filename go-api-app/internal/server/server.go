package server

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Server struct {
    router *mux.Router
}

func NewServer() *Server {
    s := &Server{
        router: mux.NewRouter(),
    }
    s.routes()
    return s
}

func (s *Server) routes() {
    s.router.HandleFunc("/health", HealthCheck).Methods("GET")
    // Add more routes here
}

func (s *Server) Start(addr string) error {
    return http.ListenAndServe(addr, s.router)
}