package clients

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes clients
func ApplyRoutes(r *gin.RouterGroup) {
	clients := r.Group("/clients")
	{
		clients.POST("", addCSVToDatabase)
	}
}
