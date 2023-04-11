package api

import (
	"AQTechincalTest-Golang/api/controller"
	"AQTechincalTest-Golang/api/route"
	"AQTechincalTest-Golang/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	r.GET("", func(context *gin.Context) { return })
	v1Api := r.Group("/v1")
	c := controller.NewController(config.ConnDB())

	route.Authentication(v1Api, c)
	route.Book(v1Api, c)

	r.Run(os.Getenv("ADDRESS"))
}
