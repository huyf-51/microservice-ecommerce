package routes

import (
	"cart/handlers"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	cart := router.Group("/cart")
	cart.Use()
	{
		cart.POST("/add-to-cart", handlers.AddToCart)
	}
}
