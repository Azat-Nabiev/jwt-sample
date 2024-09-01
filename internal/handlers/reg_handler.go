package handlers

import (
	"auth-service/internal/handlers/dto"
	"github.com/gin-gonic/gin"
)

type UserRegHandler struct {
}

func NewUserRegHandler() *UserRegHandler {
	return &UserRegHandler{}
}

func (h *UserRegHandler) Register(c *gin.Context) {
	var newUser dto.UserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

}
