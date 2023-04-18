package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"strings"
)

func (r OrderRepo) CountAllRepo(param models.OrderParam) (res int, err error) {
	var condition string
	var values []interface{}
	if param.Name != "" {
		param.Name = strings.TrimSpace(strings.ToLower(param.Name))
		condition += "AND LOWER(item_name) LIKE ?"
		values = append(values, "%"+param.Name+"%")
	}
	err = r.DB.Debug().Model(&models.Order{}).Where("deleted_at IS NULL "+condition, values...).Count(&res).Error
	return res, err
}
