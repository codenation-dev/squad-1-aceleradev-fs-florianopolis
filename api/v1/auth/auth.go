package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes auth
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("", create)
		auth.GET("", list)
		auth.GET("/:id", read)
		auth.DELETE("/:id", remove)
		auth.PATCH("/:id", update)
	}
}
