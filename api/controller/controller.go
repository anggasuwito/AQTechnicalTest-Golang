package controller

import (
	"AQTechincalTest-Golang/api/controller/c_authentication"
	"AQTechincalTest-Golang/api/controller/c_book"
	"AQTechincalTest-Golang/api/usecase"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	AuthenticationController c_authentication.AuthenticationControllerInterface
	BookController           c_book.BookControllerInterface
}

func NewController(db *gorm.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		AuthenticationController: c_authentication.NewAuthenticationController(uc),
		BookController:           c_book.NewBookController(uc),
	}
}
