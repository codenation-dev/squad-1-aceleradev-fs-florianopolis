package auth

import (
	"context"
	"crypto/sha256"
	"fmt"
	"gitlab.com/codenation-squad-1/backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
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

const userCollection = "usuarios"

func login(c *gin.Context) {
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}
	username := body.Username
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(body.Password)))

	user := getUserFromDatabase(username)
	if user.Password != password {
		c.AbortWithStatus(401)
		return
	}

	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("codenation-squad1"))
	c.JSON(200, map[string]interface{}{
		"user":  username,
		"token": tokenString,
	})
}

func create(c *gin.Context) {
	var collection = database.GetCollection(userCollection)
	var body RequestBody

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatus(400)
		return
	}

	body.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(body.Password)))
	_, err := collection.InsertOne(context.TODO(), body)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, map[string]interface{}{
		"message": "Usuário " + body.Username + " adicionado com sucesso!",
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

func getUserFromDatabase(username string) RequestBody {
	collection := database.GetCollection(userCollection)
	var res = RequestBody{}
	filter := bson.M{"username": username}
	err := collection.FindOne(database.Context, filter).Decode(&res)
	if err != nil {
		log.Println(err)
	}
	return res
}
