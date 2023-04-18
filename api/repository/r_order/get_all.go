package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"fmt"
	"strings"
)

func (r OrderRepo) GetAllRepo(offset int, limit int, param models.OrderParam) (res []models.Order, err error) {
	var condition string
	var values []interface{}
	if param.Name != "" {
		param.Name = strings.TrimSpace(strings.ToLower(param.Name))
		condition += "AND LOWER(item_name) LIKE ?"
		values = append(values, "%"+param.Name+"%")
	}
	values = append(values, limit, offset)
	fmt.Println(values)
	err = r.DB.Debug().Raw("SELECT * FROM orders WHERE deleted_at IS NULL "+condition+" ORDER BY created_at DESC LIMIT ? OFFSET ?", values...).Scan(&res).Error
	return res, err
}
