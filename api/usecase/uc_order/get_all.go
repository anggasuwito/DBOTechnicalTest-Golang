package uc_order

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/pagination"
)

func (uc OrderUC) GetAllUC(param models.OrderParam) (res []models.Order, meta pagination.Response, err error) {
	offset, limit := pagination.GetPagination(param.Request)
	res, err = uc.Repo.OrderRepo.GetAllRepo(offset, limit, param)
	if err != nil {
		return res, meta, err
	}

	total, err := uc.Repo.OrderRepo.CountAllRepo(param)
	if err != nil {
		return res, meta, err
	}

	meta = pagination.GetPaginationResponse(param.Page, limit, total)
	return res, meta, err
}
