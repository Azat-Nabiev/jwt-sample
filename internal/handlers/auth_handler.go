package handlers

import (
	"auth-service/internal/handlers/dto/authenticate"
	"auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserAuthHandler struct {
	userService *services.UserService
	logger      *zap.SugaredLogger
}

func NewUserAuthHandler(userService *services.UserService, logger *zap.SugaredLogger) *UserAuthHandler {
	return &UserAuthHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *UserAuthHandler) Authenticate(c *gin.Context) {
	var auth authenticate.UserAuthRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authResponse, err := h.userService.Authenticate(&auth)

	if err != nil {
		h.logger.Errorw("Error during registering user", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error during registering"})
		return
	}

	c.JSON(201, authResponse)
}
