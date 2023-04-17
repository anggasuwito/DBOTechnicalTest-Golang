package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
)

func (r CustomerRepo) CreateRepo(req models.Customer) (res models.Customer, err error) {
	err = r.DB.Debug().Select(
		"FullName",
		"Email",
		"Password",
		"Address",
	).Create(&req).Error
	return res, err
}
