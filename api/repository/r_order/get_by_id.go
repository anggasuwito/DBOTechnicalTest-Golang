package r_order

import "DBOTechnicalTest-Golang/api/models"

func (r OrderRepo) GetByIDRepo(id int) (res models.Order, err error) {
	err = r.DB.Debug().Where("id = ?", id).Find(&res).Error
	return res, err
}
