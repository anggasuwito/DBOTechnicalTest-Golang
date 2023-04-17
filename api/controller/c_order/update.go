package c_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c OrderController) Update(g *gin.Context) {
	var req models.Order
	g.ShouldBindJSON(&req)
	res, err := c.UC.OrderUC.UpdateUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
