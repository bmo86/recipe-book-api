package server

import (
	"errors"
	"log"
	"net/http"
	"recipe-book-api/database"
	"recipe-book-api/models"
	"recipe-book-api/repo"

	"recipe-book-api/websocket"

	"github.com/rs/cors"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Config() *models.Config
	//Hub() *websocket.Hub
}

type Broker struct {
	config *models.Config
	router *gin.Engine
	hub    *websocket.Hub
}

func (b *Broker) Config() *models.Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(config *models.Config) (*Broker, error) {

	if config.DATABASE == "" {
		return nil, errors.New("URL database empty")
	}
	if config.JWT_SECRET == "" {
		return nil, errors.New("URL database empty")
	}
	if config.PORT == "" {
		return nil, errors.New("URL database empty")
	}
	return &Broker{
		config: config,
		router: gin.New(),
		hub:    websocket.NewHub(),
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
	b.router = gin.New()
	binder(b, b.router)
	handler := cors.Default().Handler(b.router)

	repo_db, err := database.NewConnPostgres(b.config.DATABASE)
	if err != nil {
		log.Fatal(err)
	}

	go b.Hub().Run()

	repo.SetRepo(repo_db)

	log.Println("Loading server on port: ", b.config.PORT)
	if err := http.ListenAndServe(b.config.PORT, handler); err != nil {
		log.Fatal("error - Loading server on port: ", b.config.PORT)
	}

}
