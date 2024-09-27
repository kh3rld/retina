package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/mathletedev/retina/internal/db"
)

type Server struct {
	port int
	db   *db.Database
}

func NewServer(prod *bool) *http.Server {
	envPort, present := os.LookupEnv("PORT")
	if !present {
		envPort = "8080"
	}

	port, err := strconv.Atoi(envPort)
	if err != nil {
		log.Println("invalid port, using default")
	}

	allowedOrigins := []string{"*"}
	if *prod {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	s := &Server{
		port: port,
		db:   db.NewDatabase(),
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.RegisterRoutes(allowedOrigins),
	}
}
