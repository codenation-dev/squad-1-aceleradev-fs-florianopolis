package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/codenation-squad-1/backend/api"
	"gitlab.com/codenation-squad-1/backend/database"
)

//PORT port to be used
const PORT = "8080"

func main() {
	_ = database.Initialize()
	app := gin.Default()
	api.ApplyRoutes(app)    // apply api router
	_ = app.Run(":" + PORT) // listen to given port
}
