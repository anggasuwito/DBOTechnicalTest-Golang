package uc_authentication

import (
	"DBOTechnicalTest-Golang/api/models"
	"DBOTechnicalTest-Golang/helper/constants"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func (uc AuthenticationUC) LoginUC(req models.LoginRequest) (res models.LoginResponse, err error) {
	customer, err := uc.Repo.CustomerRepo.GetByEmailRepo(req.Email)
	if err != nil {
		return res, err
	}
	if customer.Email == "" {
		return res, errors.New("wrong email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password))
	if err != nil {
		return res, errors.New("wrong email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.IDKey: customer.ID,
		"exp":           jwt.NewNumericDate(time.Now().Add(constants.ExpiredTokenHour * time.Hour)),
	})

	res.Token, err = token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	res.Customer = customer
	res.Customer.Password = ""
	return res, err
}
