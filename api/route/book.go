package route

import (
	"AQTechincalTest-Golang/api/controller"
	"AQTechincalTest-Golang/api/middleware"
	"github.com/gin-gonic/gin"
)

func Book(r *gin.RouterGroup, c controller.Controller) {
	bookAPI := r.Group("/book")
	bookAPI.Use(middleware.VerifyToken())
	bookAPI.POST("", c.BookController.Create)
	bookAPI.GET("/all", c.BookController.GetAll)
	bookAPI.GET("/:id", c.BookController.GetByID)
	bookAPI.PUT("/:id", c.BookController.Update)
	bookAPI.DELETE("", c.BookController.Delete)
}
