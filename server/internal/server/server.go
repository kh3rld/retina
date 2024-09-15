package server

import (
	"flag"
	"fmt"
	"net/http"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	s := &Server{
		port: 8080,
	}

	devMode := flag.Bool("dev", false, "run in development mode")

	flag.Parse()

	allowedOrigins := []string{"*"}
	if *devMode {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.RegisterRoutes(allowedOrigins),
	}
}
