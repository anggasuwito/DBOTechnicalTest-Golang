package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	CreateRepo(req models.Order) (res models.Order, err error)
	DeleteRepo(id int) (err error)
	UpdateRepo(req models.Order) (err error)
	CountAllRepo(param models.OrderParam) (res int, err error)
	GetAllRepo(offset int, limit int, param models.OrderParam) (res []models.Order, err error)
	GetByIDRepo(id int) (res models.Order, err error)
}

func NewOrderRepo(db *gorm.DB) OrderRepoInterface {
	return &OrderRepo{DB: db}
}
