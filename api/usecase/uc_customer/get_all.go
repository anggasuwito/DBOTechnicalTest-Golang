package uc_customer

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/pagination"
)

func (uc CustomerUC) GetAllUC(param models.CustomerParam) (res []models.Customer, meta pagination.Response, err error) {
	offset, limit := pagination.GetPagination(param.Request)
	res, err = uc.Repo.CustomerRepo.GetAllRepo(offset, limit, param)
	if err != nil {
		return res, meta, err
	}

	for i := range res {
		res[i].Password = ""
	}

	total, err := uc.Repo.CustomerRepo.CountAllRepo(param)
	if err != nil {
		return res, meta, err
	}

	meta = pagination.GetPaginationResponse(param.Page, limit, total)
	return res, meta, err
}
