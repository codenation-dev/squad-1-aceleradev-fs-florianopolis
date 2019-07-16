package v1

import (
	"gitlab.com/codenation-squad-1/backend/api/v1/auth"
	"gitlab.com/codenation-squad-1/backend/api/v1/clients"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		auth.ApplyRoutes(v1)
		clients.ApplyRoutes(v1)
	}
}
