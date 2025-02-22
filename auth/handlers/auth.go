package handlers

import (
	"auth/configs"
	"auth/models"
	"auth/types"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var requestUser types.User
	var user models.User
	if c.ShouldBind(&requestUser) == nil {
		configs.DB.Where("username = ?", requestUser.Username).First(&user)
		if user.ID != 0 {
			c.JSON(403, gin.H{
				"message": "user exist",
			})
			return
		}
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(requestUser.Password), 14)
		user = models.User{Username: requestUser.Username, Password: string(hashPassword)}
		configs.DB.Create(&user)
	}
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func Login(c *gin.Context) {
	var requestUser types.User
	var user models.User
	if c.ShouldBind(&requestUser) == nil {
		configs.DB.Where("username = ?", requestUser.Username).First(&user)
		if user.ID == 0 {
			c.JSON(403, gin.H{
				"message": "incorrect username or password",
			})
			return
		}
		fmt.Println("username>>>", user.Username)
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password))
		if err != nil {
			c.JSON(403, gin.H{
				"message": "incorrect username or password",
			})
			return
		}
	}
	pkpem, _ := os.ReadFile("./key/private-key.pem")
	pk, _ := jwt.ParseECPrivateKeyFromPEM(pkpem)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{ExpiresAt: time.Now().Unix() + 15000})
	accessToken, _ := token.SignedString(pk)

	c.JSON(200, gin.H{
		"accessToken": accessToken,
	})
}
