package c_book

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c BookController) Create(g *gin.Context) {
	var req models.BookDetails
	g.ShouldBindJSON(&req)
	err := c.UC.BookUC.CreateUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  nil,
		Meta:  nil,
		Error: err,
	})
}
