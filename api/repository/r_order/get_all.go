package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"strings"
)

func (r OrderRepo) GetAllRepo(offset int, limit int, param models.OrderParam) (res []models.Order, err error) {
	param.Name = strings.TrimSpace(strings.ToLower(param.Name))
	err = r.DB.Debug().Where("deleted_at IS NULL").Order("created_at DESC").Limit(limit).Offset(offset).Find(&res).Error
	return res, err
}
