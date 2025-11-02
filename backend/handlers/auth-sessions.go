package handlers

import (
	"IdeaWeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
This handler will register & login users and will authenticate via 2 following methods:
1) Session-based with cookies
2) JWT
*/

type AuthSessionsHandler struct {
	ss *services.SessionService
}

func NewAuthSessionsHandler (ss *services.SessionService) {
	return &AuthSessionsHandler{
		ss: ss,
	}
}

func Login(c *gin.Context) {
	
}

func Register(c *gin.Context) {
	var registerReq models.RegisterRequest

	if err := c.ShouldBinJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid registration format"
		})
		return
	}


}
