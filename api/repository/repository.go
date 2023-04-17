package repository

import (
	"DBOTechnicalTest-Golang/api/repository/r_customer"
	"DBOTechnicalTest-Golang/api/repository/r_order"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	CustomerRepo r_customer.CustomerRepoInterface
	OrderRepo    r_order.OrderRepoInterface
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		CustomerRepo: r_customer.NewCustomerRepo(db),
		OrderRepo:    r_order.NewOrderRepo(db),
	}
}
