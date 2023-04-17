package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"strings"
)

func (r CustomerRepo) GetByEmailRepo(email string) (res models.Customer, err error) {
	email = strings.TrimSpace(strings.ToLower(email))
	err = r.DB.Debug().Where("LOWER(email) = ?", email).Find(&res).Error
	return res, err
}
