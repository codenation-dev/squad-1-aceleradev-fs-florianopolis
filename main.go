package main

import (
	"github.com/gin-gonic/gin"
)

//PORT port to be used
const PORT = "8080"

func main() {

	app := gin.Default()
	app.Run(":" + PORT) // listen to given port
}
