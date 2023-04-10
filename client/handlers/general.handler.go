package handlers

import (
	"net/http"
	"recipe-book-api/server"

	"github.com/gin-gonic/gin"
)

func HandlerHome(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, is me server! ---- Im Robot xd"})
	}
}
