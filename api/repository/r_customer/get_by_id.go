package r_customer

import "DBOTechnicalTest-Golang/api/models"

func (r CustomerRepo) GetByIDRepo(id int) (res models.Customer, err error) {
	err = r.DB.Debug().Raw("SELECT * FROM customer WHERE deleted_at IS NULL AND id = ?", id).Scan(&res).Error
	return res, err
}
