package main

import (
	"log"
	"os"
	"recipe-book-api/client/handlers"
	"recipe-book-api/models"
	"recipe-book-api/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	JWT := os.Getenv("JWT")
	DB := os.Getenv("DB")

	s, err := server.NewServer(&models.Config{PORT: PORT, JWT_SECRET: JWT, DATABASE: DB})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(Router)
}

func Router(s server.Server, r *gin.Engine) {
	//api := r.Group("/api/v1")
	r.GET("/home", handlers.HandlerHome(s))

}
