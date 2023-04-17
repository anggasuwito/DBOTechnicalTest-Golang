package usecase

import (
	"DBOTechnicalTest-Golang/api/repository"
	"DBOTechnicalTest-Golang/api/usecase/uc_authentication"
	"DBOTechnicalTest-Golang/api/usecase/uc_customer"
	"DBOTechnicalTest-Golang/api/usecase/uc_order"
	"github.com/jinzhu/gorm"
)

type UseCase struct {
	AuthenticationUC uc_authentication.AuthenticationUCInterface
	CustomerUC       uc_customer.CustomerUCInterface
	OrderUC          uc_order.OrderUCInterface
}

func NewUseCase(db *gorm.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		AuthenticationUC: uc_authentication.NewAuthenticationUC(repo),
		CustomerUC:       uc_customer.NewCustomerUC(repo),
		OrderUC:          uc_order.NewOrderUC(repo),
	}
}
