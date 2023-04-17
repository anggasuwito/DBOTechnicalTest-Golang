package c_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c OrderController) Create(g *gin.Context) {
	var req models.Order
	g.ShouldBindJSON(&req)
	res, err := c.UC.OrderUC.CreateUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
