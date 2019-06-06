package main

import (
	"gitlab.com/codenation-squad-1/backend/api"

	"github.com/gin-gonic/gin"
)

//PORT port to be used
const PORT = "8080"

func main() {

	app := gin.Default()
	api.ApplyRoutes(app) // apply api router
	app.Run(":" + PORT)  // listen to given port
}
