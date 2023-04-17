package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
)

func (r OrderRepo) DeleteRepo(id int) (err error) {
	return r.DB.Debug().Where("id = ?", id).Delete(&models.Order{}).Error
}
