package v1

import (
	"aceleradev/backend/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		auth.ApplyRoutes(v1)
	}
}
