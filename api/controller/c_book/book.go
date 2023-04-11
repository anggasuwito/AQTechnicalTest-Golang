package c_book

import (
	"AQTechincalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	UC usecase.UseCase
}

type BookControllerInterface interface {
	Create(g *gin.Context)
	GetAll(g *gin.Context)
	GetByID(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

func NewBookController(uc usecase.UseCase) BookControllerInterface {
	return &BookController{UC: uc}
}
