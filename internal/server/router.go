package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jason2071/pets/internal/handler"
)

// NewRouter builds the Gin engine and registers all routes.
func NewRouter(petHandler *handler.PetHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")
	petHandler.Register(api)

	return r
}
