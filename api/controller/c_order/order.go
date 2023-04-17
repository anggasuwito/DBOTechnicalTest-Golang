package c_order

import (
	"DBOTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	UC usecase.UseCase
}

type OrderControllerInterface interface {
	Create(g *gin.Context)
	GetAll(g *gin.Context)
	GetByID(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}

func NewOrderController(uc usecase.UseCase) OrderControllerInterface {
	return &OrderController{UC: uc}
}
