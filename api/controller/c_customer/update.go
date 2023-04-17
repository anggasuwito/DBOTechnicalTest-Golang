package c_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c CustomerController) Update(g *gin.Context) {
	var req models.Customer
	g.ShouldBindJSON(&req)
	res, err := c.UC.CustomerUC.UpdateUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
