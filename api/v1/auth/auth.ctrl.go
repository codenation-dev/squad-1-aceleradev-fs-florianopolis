package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func login(c *gin.Context) {
	var body RequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	user := body.Username
	password := body.Password

	if user != "test" && password != "test" {
		c.AbortWithStatus(401)
	}

	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("codenation-squad1"))
	c.JSON(200, map[string]interface{}{
		"user":  user,
		"token": tokenString,
	})
}

func list(c *gin.Context) {
	c.Status(204)
}

func read(c *gin.Context) {
	c.Status(200)
}

func remove(c *gin.Context) {
	c.Status(204)
}

func update(c *gin.Context) {
	c.Status(200)
}
