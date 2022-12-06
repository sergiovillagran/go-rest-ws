package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sergiovillagran/rest-ws/database"
	"github.com/sergiovillagran/rest-ws/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
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
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("Secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("Db url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal("Impossible connect to the db")
	}

	repository.SetRepository(repo)
	log.Println("Starting server on port", b.config.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", b.config.Port), b.router); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}
