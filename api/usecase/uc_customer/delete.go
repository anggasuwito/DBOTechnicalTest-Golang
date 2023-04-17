package uc_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
	"strconv"
)

func (uc CustomerUC) DeleteUC(id string) (res models.Customer, err error) {
	numID, _ := strconv.Atoi(id)
	res, err = uc.Repo.CustomerRepo.GetByIDRepo(numID)
	if err != nil {
		return res, err
	}
	if res.ID == 0 {
		return res, errors.New("customer not found")
	}

	err = uc.Repo.CustomerRepo.DeleteRepo(numID)
	return res, err
}
