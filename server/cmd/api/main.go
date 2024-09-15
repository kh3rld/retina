package main

import (
	"log"

	"github.com/mathletedev/retina/internal/server"
)

func main() {
	s := server.NewServer()

	log.Println("server started!")
	log.Fatal(s.ListenAndServe())
}
