package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jason2071/pets/internal/auth"
	"github.com/jason2071/pets/internal/domain"
	"github.com/jason2071/pets/internal/service"
	"gorm.io/gorm"
)

// PetHandler handles HTTP requests for pets.
type PetHandler struct {
	svc *service.PetService
}

// NewPetHandler constructs a PetHandler.
func NewPetHandler(svc *service.PetService) *PetHandler {
	return &PetHandler{svc: svc}
}

type petRequest struct {
	Name    string `json:"name" binding:"required"`
	Species string `json:"species" binding:"required"`
	Breed   string `json:"breed"`
	Age     int    `json:"age"`
}

// Register wires pet routes onto the given router group.
func (h *PetHandler) Register(rg *gin.RouterGroup) {
	pets := rg.Group("/pets")
	{
		pets.POST("", h.Create)
		pets.GET("", h.List)
		pets.GET("/:id", h.Get)
		pets.PUT("/:id", h.Update)
		pets.DELETE("/:id", h.Delete)
	}
}

func (h *PetHandler) Create(c *gin.Context) {
	var req petRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ownerID, ok := accountIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		return
	}
	p := &domain.Pet{OwnerID: &ownerID, Name: req.Name, Species: req.Species, Breed: req.Breed, Age: req.Age}
	if err := h.svc.Create(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *PetHandler) List(c *gin.Context) {
	ownerID, ok := accountIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		return
	}
	pets, err := h.svc.List(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pets)
}

func (h *PetHandler) Get(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ownerID, ok := accountIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		return
	}
	p, err := h.svc.Get(id, ownerID)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *PetHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req petRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ownerID, ok := accountIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		return
	}
	in := &domain.Pet{Name: req.Name, Species: req.Species, Breed: req.Breed, Age: req.Age}
	p, err := h.svc.Update(id, ownerID, in)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *PetHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	ownerID, ok := accountIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token subject"})
		return
	}
	if err := h.svc.Delete(id, ownerID); err != nil {
		h.respondError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *PetHandler) respondError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "pet not found"})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// accountIDFromToken extracts the authenticated account id from JWT claims
// (the "sub" claim) set by the auth middleware.
func accountIDFromToken(c *gin.Context) (uint, bool) {
	claims, ok := auth.ClaimsFrom(c)
	if !ok {
		return 0, false
	}
	id, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(id), true
}
