package handlers

import (
	"IdeaWeb/models"
	"IdeaWeb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	sessionService *services.SessionService
}

func NewSessionHandler(sessionService *services.SessionService) *SessionHandler {
	return &SessionHandler{sessionService: sessionService}
}

func (h *SessionHandler) Register(c *gin.Context) {
	var registerReq *models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid input"},
		)
		return
	}

	err := h.sessionService.Register(registerReq)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"error": "registration successful"},
	)
}

func (h *SessionHandler) Login(c *gin.Context) {
	var loginReq *models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid input"},
		)
		return
	}

	sessionID, err := h.sessionService.Login(loginReq)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "login failed: " + err.Error()},
		)
		return
	}

	c.SetCookie(
		"session_id",
		sessionID,
		3600,
		"/",
		"localhost",
		true,
		true,
	)

	c.JSON(
		http.StatusOK,
		gin.H{"message": "login successful"},
	)
}
