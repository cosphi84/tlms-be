package handlers

import (
	"net/http"
	"tlms/internal/dto"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthenticateHandler struct {
	authService services.AuthenticateService
}

func NewAuthenticateHandler(authService services.AuthenticateService) *AuthenticateHandler {
	return &AuthenticateHandler{
		authService: authService,
	}
}

func (h *AuthenticateHandler) Authenticate(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ipAddress := c.ClientIP()

	authRetrn, err := h.authService.Authenticate(&req, ipAddress)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"access_token":  authRetrn.AccessToken,
			"refresh_token": authRetrn.RefreshToken,
			"user":          authRetrn.User,
		},
	)
}

func (h *AuthenticateHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	authRetrn, err := h.authService.RefreshToken(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"access_token":  authRetrn.AccessToken,
			"refresh_token": authRetrn.RefreshToken,
			"user":          authRetrn.User,
		},
	)
}
