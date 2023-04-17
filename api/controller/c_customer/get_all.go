package c_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c CustomerController) GetAll(g *gin.Context) {
	var req models.CustomerParam
	g.ShouldBindQuery(&req)
	res, meta, err := c.UC.CustomerUC.GetAllUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  meta,
		Error: err,
	})
}
