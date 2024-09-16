package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/mathletedev/retina/internal/auth"
	"github.com/mathletedev/retina/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	prod := flag.Bool("dev", false, "run in development mode")

	flag.Parse()

	*prod = !*prod

	auth.NewAuth(prod)

	s := server.NewServer(prod)

	log.Println("server started!")
	log.Fatal(s.ListenAndServe())
}
