package uc_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
)

func (uc CustomerUC) UpdateUC(req models.Customer) (res models.Customer, err error) {
	customer, err := uc.Repo.CustomerRepo.GetByIDRepo(req.ID)
	if err != nil {
		return res, err
	}
	if customer.ID == 0 {
		return res, errors.New("customer not found")
	}

	existingCustomer, err := uc.Repo.CustomerRepo.GetByEmailRepo(req.Email)
	if err != nil {
		return res, err
	}
	if existingCustomer.ID != 0 {
		return res, errors.New("this email has been registered")
	}

	err = uc.Repo.CustomerRepo.UpdateRepo(req)
	if err != nil {
		return res, err
	}

	return uc.Repo.CustomerRepo.GetByIDRepo(req.ID)
}
