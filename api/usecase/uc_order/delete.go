package uc_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
	"strconv"
)

func (uc OrderUC) DeleteUC(id string) (res models.Order, err error) {
	numID, _ := strconv.Atoi(id)
	res, err = uc.Repo.OrderRepo.GetByIDRepo(numID)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("order not found")
	}

	err = uc.Repo.OrderRepo.DeleteRepo(numID)
	return res, err
}
