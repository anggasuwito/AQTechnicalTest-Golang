package c_book

import (
	"AQTechincalTest-Golang/helper/pagination"
	"AQTechincalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c BookController) GetAll(g *gin.Context) {
	var req pagination.Request
	g.ShouldBindQuery(&req)
	res, meta, err := c.UC.BookUC.GetAllUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  meta,
		Error: err,
	})
}
