package r_order

import "DBOTechnicalTest-Golang/api/models"

func (r OrderRepo) CountAllRepo() (res int, err error) {
	err = r.DB.Debug().Model(&models.Order{}).Count(&res).Error
	return res, err
}
