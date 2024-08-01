package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cristiangar0398/leal-rewards/internal/adapters/inbound/handlers"
	"github.com/cristiangar0398/leal-rewards/internal/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		BatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
