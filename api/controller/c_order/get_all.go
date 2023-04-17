package c_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c OrderController) GetAll(g *gin.Context) {
	var req models.OrderParam
	g.ShouldBindQuery(&req)
	res, meta, err := c.UC.OrderUC.GetAllUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  meta,
		Error: err,
	})
}
