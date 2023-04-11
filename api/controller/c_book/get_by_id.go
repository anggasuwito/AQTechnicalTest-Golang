package c_book

import (
	"AQTechincalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c BookController) GetByID(g *gin.Context) {
	id := g.Param("id")
	res, err := c.UC.BookUC.GetByIDUC(id)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
