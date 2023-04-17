package c_customer

import (
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c CustomerController) Delete(g *gin.Context) {
	id := g.Param("id")
	res, err := c.UC.CustomerUC.DeleteUC(id)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
