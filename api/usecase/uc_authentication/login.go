package uc_authentication

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/helper/constants"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func (uc AuthenticationUC) LoginUC(req models.LoginRequest) (res string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.UsernameKey: req.Username,
		"exp":                 jwt.NewNumericDate(time.Now().Add(constants.ExpiredTokenHour * time.Hour)),
	})

	res, err = token.SignedString([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	return res, err
}
