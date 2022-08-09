package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login-with-email", loginWithEmail)
	}
}

func loginWithEmail(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"email": "test@gmail.com",
	})
}
