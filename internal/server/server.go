package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/cristiangar0398/leal-rewards/database"
	"github.com/cristiangar0398/leal-rewards/internal/adapters/outbound/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	BatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.BatabaseUrl == "" {
		return nil, errors.New("Data Base is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(bainder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	bainder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.BatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)
	port := b.Config().Port

	log.Println(">>> >>> >>> 🚀 El servidor está despegando en el puerto", port, ">>> >>> >>>")
	if err := http.ListenAndServe(":"+b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
