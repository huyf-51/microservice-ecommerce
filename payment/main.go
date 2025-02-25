package main

import (
	"fmt"
	"payment/configs"
	"payment/grpc"

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
	grpc.ConnectGrpcServer()
	router.Run(":3003")
}
