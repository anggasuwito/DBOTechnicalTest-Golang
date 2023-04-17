package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
)

func (r CustomerRepo) DeleteRepo(id int) (err error) {
	return r.DB.Debug().Where("id = ?", id).Delete(&models.Customer{}).Error
}
