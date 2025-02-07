package main

import (
	"cart/configs"
	"cart/handlers"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	godotenv.Load()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	_, err := configs.ConnectDB()
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})
	e := rdb.Set(context.Background(), "name", "huy", 3600000).Err()
	if e != nil {
		fmt.Println(e)
	}
	if err != nil {
		fmt.Println("connect db error")
	}
	handlers.Api(router)
	router.Run(":3002")
}
