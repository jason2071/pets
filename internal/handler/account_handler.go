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

type registerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AccountHandler) Register(rg *gin.RouterGroup) {
	account := rg.Group("/account")
	{
		account.POST("register", h.Create)
		account.POST("login", h.Login)
	}
}

func (h *AccountHandler) Create(c *gin.Context) {
	var req registerRequest
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

func (h *AccountHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.account.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": service.ErrInvalidCredentials.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
