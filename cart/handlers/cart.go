package handlers

import (
	"cart/configs"
	"cart/models"
	"cart/types"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var requestCart *types.Cart

	if c.ShouldBind(&requestCart) == nil {
		fmt.Println(requestCart)
		configs.Rdb.HSet(context.Background(), "cart:"+fmt.Sprint(requestCart.UserId), "product:"+fmt.Sprint(requestCart.Item.ProductId), requestCart.Item.Quantity)
		cart := models.Cart{UserId: requestCart.UserId}
		configs.DB.Create(&cart)
		item := models.Item{ProductId: requestCart.Item.ProductId, Quantity: requestCart.Item.Quantity, CartID: cart.ID}
		configs.DB.Create(&item)
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}
