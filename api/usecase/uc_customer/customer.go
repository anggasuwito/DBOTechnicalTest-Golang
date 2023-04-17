package uc_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/api/repository"
	"DBOTechnicalTest-Golang/helper/pagination"
)

type CustomerUC struct {
	Repo repository.Repository
}

type CustomerUCInterface interface {
	CreateUC(req models.Customer) (res models.Customer, err error)
	GetAllUC(param models.CustomerParam) (res []models.Customer, meta pagination.Response, err error)
	GetByIDUC(id string) (res models.Customer, err error)
	UpdateUC(req models.Customer) (res models.Customer, err error)
	DeleteUC(id string) (res models.Customer, err error)
}

func NewCustomerUC(repo repository.Repository) CustomerUCInterface {
	return &CustomerUC{Repo: repo}
}
