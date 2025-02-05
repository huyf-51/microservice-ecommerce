package main

import (
	"auth/configs"
	"auth/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	err := configs.ConnectDB()
	if err != nil {
		fmt.Println("connect db error")
	}
	handlers.Api(router)
	router.Run(":3001")
}
