package r_customer

import "DBOTechnicalTest-Golang/api/models"

func (r CustomerRepo) CountAllRepo() (res int, err error) {
	err = r.DB.Debug().Model(&models.Customer{}).Count(&res).Error
	return res, err
}
