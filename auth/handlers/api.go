package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/auth/login", func(c *gin.Context) {
		fmt.Println("hello auth route")
		c.JSON(200, gin.H{
			"hello": "huy",
		})
	})
}
