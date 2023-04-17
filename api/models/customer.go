package models

import "DBOTechnicalTest-Golang/helper/pagination"

func (Customer) TableName() string {
	return "customer"
}

type Customer struct {
	ID        int    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime"`
	DeletedAt string `json:"deleted_at" gorm:"type:datetime"`
}

type CustomerParam struct {
	Name string `form:"name"`
	pagination.Request
}
