package route

import (
	"AQTechincalTest-Golang/api/controller"
	"github.com/gin-gonic/gin"
)

func Authentication(r *gin.RouterGroup, c controller.Controller) {
	auth := r.Group("/auth")
	auth.POST("/login", c.AuthenticationController.Login)
}
