package middleware

import (
	"IdeaWeb/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareModel struct {
	sessionService *services.SessionService
}

func NewAuthMiddleware(ss *services.SessionService) *AuthMiddlewareModel {
	return &AuthMiddlewareModel{
		sessionService: ss,
	}
}

func (m *AuthMiddlewareModel) AuthMiddleware(c *gin.Context) {
	// check cookies
	sID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "cookies not found"},
		)
	}

	sessionID, err := strconv.ParseUint(sID, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "unable to parse session id"},
		)
	}

	_, err = m.sessionService.ValidateSession(uint(sessionID))
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "failed to validate session"},
		)
	}

	c.Next()
}
