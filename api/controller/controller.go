package controller

import (
	"DBOTechnicalTest-Golang/api/controller/c_authentication"
	"DBOTechnicalTest-Golang/api/controller/c_customer"
	"DBOTechnicalTest-Golang/api/controller/c_order"
	"DBOTechnicalTest-Golang/api/usecase"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	AuthenticationController c_authentication.AuthenticationControllerInterface
	CustomerController       c_customer.CustomerControllerInterface
	OrderController          c_order.OrderControllerInterface
}

func NewController(db *gorm.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		AuthenticationController: c_authentication.NewAuthenticationController(uc),
		CustomerController:       c_customer.NewCustomerController(uc),
		OrderController:          c_order.NewOrderController(uc),
	}
}
