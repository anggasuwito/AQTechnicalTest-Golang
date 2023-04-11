package c_book

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c BookController) Update(g *gin.Context) {
	var req models.Book
	g.ShouldBindJSON(&req)
	err := c.UC.BookUC.UpdateUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  nil,
		Meta:  nil,
		Error: err,
	})
}
