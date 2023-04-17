package uc_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
)

func (uc OrderUC) CreateUC(req models.Order) (res models.Order, err error) {
	if req.CustomerID == 0 {
		return res, errors.New("please input customer id")
	} else if req.ItemName == "" {
		return res, errors.New("please input item name")
	} else if req.Quantity == 0 {
		return res, errors.New("please input item quantity")
	} else if req.Price == 0 {
		return res, errors.New("please input item price")
	}

	return uc.Repo.OrderRepo.CreateRepo(req)
}
