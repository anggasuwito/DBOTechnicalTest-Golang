package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"strings"
)

func (r CustomerRepo) CountAllRepo(param models.CustomerParam) (res int, err error) {
	var condition string
	var values []interface{}
	if param.Name != "" {
		param.Name = strings.TrimSpace(strings.ToLower(param.Name))
		condition += "AND LOWER(full_name) LIKE ?"
		values = append(values, "%"+param.Name+"%")
	}
	err = r.DB.Debug().Model(&models.Customer{}).Where("deleted_at IS NULL "+condition, values...).Count(&res).Error
	return res, err
}
