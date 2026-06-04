package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jason2071/pets/internal/auth"
	"github.com/jason2071/pets/internal/handler"
)

// NewRouter builds the Gin engine and registers all routes.
func NewRouter(petHandler *handler.PetHandler, accHandler *handler.AccountHandler, tokens *auth.TokenManager) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")

	// Public: register + login.
	accHandler.Register(api)

	// Protected: everything below requires a valid bearer token.
	secured := api.Group("")
	secured.Use(tokens.RequireToken())
	petHandler.Register(secured)

	return r
}
