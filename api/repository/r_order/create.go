package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
)

func (r OrderRepo) CreateRepo(req models.Order) (res models.Order, err error) {
	err = r.DB.Debug().Select(
		"CustomerID",
		"ItemName",
		"Quantity",
		"Price",
	).Create(&req).Error
	return res, err
}
