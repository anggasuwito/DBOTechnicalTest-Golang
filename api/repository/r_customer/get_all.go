package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"strings"
)

func (r CustomerRepo) GetAllRepo(offset int, limit int, param models.CustomerParam) (res []models.Customer, err error) {
	var condition string
	var values []interface{}
	if param.Name != "" {
		param.Name = strings.TrimSpace(strings.ToLower(param.Name))
		condition += "AND LOWER(full_name) LIKE ?"
		values = append(values, "%"+param.Name+"%")
	}
	values = append(values, limit, offset)

	err = r.DB.Debug().Raw("SELECT * FROM customer WHERE deleted_at IS NULL "+condition+" LIMIT ? OFFSET ?", values...).Scan(&res).Error
	return res, err
}
