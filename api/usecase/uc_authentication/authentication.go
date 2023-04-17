package uc_authentication

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/api/repository"
)

type AuthenticationUC struct {
	Repo repository.Repository
}

type AuthenticationUCInterface interface {
	LoginUC(req models.LoginRequest) (res models.LoginResponse, err error)
}

func NewAuthenticationUC(repo repository.Repository) AuthenticationUCInterface {
	return &AuthenticationUC{Repo: repo}
}
