package c_order

import (
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c OrderController) Delete(g *gin.Context) {
	id := g.Param("id")
	res, err := c.UC.OrderUC.DeleteUC(id)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
