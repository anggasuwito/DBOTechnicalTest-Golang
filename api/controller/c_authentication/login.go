package c_authentication

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c AuthenticationController) Login(g *gin.Context) {
	var req models.LoginRequest
	g.ShouldBindJSON(&req)
	res, err := c.UC.AuthenticationUC.LoginUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
