package r_customer

import "DBOTechnicalTest-Golang/api/models"

func (r CustomerRepo) GetByIDRepo(id int) (res models.Customer, err error) {
	err = r.DB.Debug().Where("id = ?", id).Find(&res).Error
	return res, err
}
