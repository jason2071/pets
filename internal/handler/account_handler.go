package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jason2071/pets/internal/domain"
	"github.com/jason2071/pets/internal/service"
)

type AccountHandler struct {
	account *service.AccountService
}

func NewAccountHandler(account *service.AccountService) *AccountHandler {
	return &AccountHandler{account: account}
}

type accountRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func (h *AccountHandler) Register(rg *gin.RouterGroup) {
	account := rg.Group("/account")
	{
		account.POST("register", h.Create)
	}
}

func (h *AccountHandler) Create(c *gin.Context) {
	var req accountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc := &domain.Account{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	if err := h.account.Create(acc); err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			c.JSON(http.StatusConflict, gin.H{"error": service.ErrEmailExists.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Register Complete")
}
