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

func LoginHandler(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid login format",
		})
		return
	}

	// Else, authenticate the user and create a session for them
}
