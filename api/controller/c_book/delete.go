package c_book

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c BookController) Delete(g *gin.Context) {
	var req models.BookListID
	g.ShouldBindJSON(&req)
	err := c.UC.BookUC.DeleteUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  nil,
		Meta:  nil,
		Error: err,
	})
}
