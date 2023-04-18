package r_order

import "DBOTechnicalTest-Golang/api/models"

func (r OrderRepo) GetByIDRepo(id int) (res models.Order, err error) {
	err = r.DB.Debug().Raw("SELECT * FROM orders WHERE deleted_at IS NULL AND id = ?", id).Scan(&res).Error
	return res, err
}
