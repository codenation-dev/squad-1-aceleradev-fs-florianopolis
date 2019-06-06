package api

import (
	apiv1 "aceleradev/backend/api/v1"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes v1
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
