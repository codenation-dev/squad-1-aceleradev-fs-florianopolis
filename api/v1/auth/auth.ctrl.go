package auth

import (
	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	c.Status(200)
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
