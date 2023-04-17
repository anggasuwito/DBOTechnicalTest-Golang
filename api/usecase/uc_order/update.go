package uc_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
)

func (uc OrderUC) UpdateUC(req models.Order) (res models.Order, err error) {
	order, err := uc.Repo.OrderRepo.GetByIDRepo(req.ID)
	if err != nil {
		return res, err
	}
	if order.ID == 0 {
		return res, errors.New("order not found")
	}

	err = uc.Repo.OrderRepo.UpdateRepo(req)
	if err != nil {
		return res, err
	}

	return uc.Repo.OrderRepo.GetByIDRepo(req.ID)
}
