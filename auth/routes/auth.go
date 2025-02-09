package routes

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.Use()
	{
		auth.POST("/signup", handlers.SignUp)
		auth.POST("/login", handlers.Login)
	}

}
