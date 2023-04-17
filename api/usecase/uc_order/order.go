package uc_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/api/repository"
	"DBOTechnicalTest-Golang/helper/pagination"
)

type OrderUC struct {
	Repo repository.Repository
}

type OrderUCInterface interface {
	CreateUC(req models.Order) (res models.Order, err error)
	GetAllUC(param models.OrderParam) (res []models.Order, meta pagination.Response, err error)
	GetByIDUC(id string) (res models.Order, err error)
	UpdateUC(req models.Order) (res models.Order, err error)
	DeleteUC(id string) (res models.Order, err error)
}

func NewOrderUC(repo repository.Repository) OrderUCInterface {
	return &OrderUC{Repo: repo}
}
