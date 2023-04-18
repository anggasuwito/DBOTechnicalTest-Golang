package uc_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func (uc CustomerUC) CreateUC(req models.Customer) (res models.Customer, err error) {
	if req.Email == "" {
		return res, errors.New("please input email")
	} else if req.Password == "" {
		return res, errors.New("please input password")
	}

	customer, err := uc.Repo.CustomerRepo.GetByEmailRepo(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return res, err
	}
	if customer.ID != 0 {
		return res, errors.New("this email has been registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}
	req.Password = string(hashedPassword)
	return uc.Repo.CustomerRepo.CreateRepo(req)
}
