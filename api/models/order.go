package models

import "DBOTechnicalTest-Golang/helper/pagination"

func (Order) TableName() string {
	return "order"
}

type Order struct {
	ID         int     `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	CustomerID int     `json:"customer_id"`
	ItemName   string  `json:"item_name"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	CreatedAt  string  `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	UpdatedAt  string  `json:"updated_at" gorm:"type:datetime"`
	DeletedAt  string  `json:"deleted_at" gorm:"type:datetime"`
}

type OrderParam struct {
	Name string `form:"name"`
	pagination.Request
}
