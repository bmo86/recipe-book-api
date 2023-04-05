package server

import (
	"errors"
	"recipe-book-api/models"

	"recipe-book-api/websocket"

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

// func (b *Broker) Start(binder func(s Server, r *gin.Engine)) {
// 	b.router = gin.New()
// 	binder(b, b.router)
// 	handler := cors.Default().Handler(b.router)

// 	go b.Hub().Run()

// }
