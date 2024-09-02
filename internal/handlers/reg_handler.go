package handlers

import (
	"auth-service/internal/handlers/dto/reg"
	"auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserRegHandler struct {
	userService *services.UserService
	logger      *zap.SugaredLogger
}

func NewUserRegHandler(userService *services.UserService, logger *zap.SugaredLogger) *UserRegHandler {
	return &UserRegHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserRegHandler) Register(c *gin.Context) {
	var newUser reg.UserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	regResponse, err := h.userService.Register(&newUser)

	if err != nil {
		h.logger.Errorw("Error during registering user", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error during registering"})
		return
	}

	c.JSON(201, regResponse)
}
