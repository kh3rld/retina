package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
}

func NewServer(prod *bool) *http.Server {
	s := &Server{
		port: 8080,
	}

	allowedOrigins := []string{"*"}
	if *prod {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.RegisterRoutes(allowedOrigins),
	}
}
