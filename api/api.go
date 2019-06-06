package api

import (
	apiv1 "gitlab.com/codenation-squad-1/backend/api/v1"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes v1
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiv1.ApplyRoutes(api)
	}
}
