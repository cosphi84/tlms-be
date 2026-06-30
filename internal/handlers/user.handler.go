package handlers

import (
	"net/http"
	"tlms/internal/dto"
	"tlms/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(
	userService services.UserService,
) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		c.Abort()
		return
	}

	err := handler.userService.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})

}
