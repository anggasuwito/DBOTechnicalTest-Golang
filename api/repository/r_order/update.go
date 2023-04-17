package r_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"time"
)

func (r OrderRepo) UpdateRepo(req models.Order) (err error) {
	return r.DB.Debug().Model(&models.Order{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"customer_id": req.CustomerID,
		"item_name":   req.ItemName,
		"quantity":    req.Quantity,
		"price":       req.Price,
		"updated_at":  time.Now().Format(time.RFC3339),
	}).Error
}
