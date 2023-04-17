package r_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"time"
)

func (r CustomerRepo) UpdateRepo(req models.Customer) (err error) {
	return r.DB.Debug().Model(&models.Customer{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"full_name":  req.FullName,
		"email":      req.Email,
		"address":    req.Address,
		"updated_at": time.Now().Format(time.RFC3339),
	}).Error
}
