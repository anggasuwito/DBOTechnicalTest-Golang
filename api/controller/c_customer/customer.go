package c_customer

import (
	"DBOTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	UC usecase.UseCase
}

type CustomerControllerInterface interface {
	Create(g *gin.Context)
	GetAll(g *gin.Context)
	GetByID(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

func NewCustomerController(uc usecase.UseCase) CustomerControllerInterface {
	return &CustomerController{UC: uc}
}
