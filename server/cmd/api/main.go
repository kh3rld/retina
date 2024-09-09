package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	devMode := flag.Bool("dev", false, "run in development mode")

	flag.Parse()

	allowedOrigins := []string{"*"}
	if *devMode {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	c := cors.New(cors.Options{
		AllowedOrigins: allowedOrigins,
	})
	handler := c.Handler(http.DefaultServeMux)

	http.HandleFunc("/api/hello", handleGetHello)

	log.Println("server started!")
	http.ListenAndServe(":8080", handler)
}

func handleGetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
