package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"github.com/jinzhu/gorm"
)

type CustomerRepo struct {
	DB *gorm.DB
}

type CustomerRepoInterface interface {
	CreateRepo(req models.Customer) (res models.Customer, err error)
	DeleteRepo(id int) (err error)
	UpdateRepo(req models.Customer) (err error)
	CountAllRepo() (res int, err error)
	GetAllRepo(offset int, limit int, param models.CustomerParam) (res []models.Customer, err error)
	GetByIDRepo(id int) (res models.Customer, err error)
	GetByEmailRepo(email string) (res models.Customer, err error)
}

func NewCustomerRepo(db *gorm.DB) CustomerRepoInterface {
	return &CustomerRepo{DB: db}
}
