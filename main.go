package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/codenation-squad-1/backend/api"
	"gitlab.com/codenation-squad-1/backend/database"
	"log"
	"os"
)

//PORT port to be used
var PORT string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT = os.Getenv("SERVER_PORT")
	_ = database.Initialize()
	app := gin.Default()
	api.ApplyRoutes(app)    // apply api router
	_ = app.Run(":" + PORT) // listen to given port
}
