package usecase

import (
	"AQTechincalTest-Golang/api/repository"
	"AQTechincalTest-Golang/api/usecase/uc_authentication"
	"AQTechincalTest-Golang/api/usecase/uc_book"
	"github.com/jinzhu/gorm"
)

type UseCase struct {
	AuthenticationUC uc_authentication.AuthenticationUCInterface
	BookUC           uc_book.BookUCInterface
}

func NewUseCase(db *gorm.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		AuthenticationUC: uc_authentication.NewAuthenticationUC(repo),
		BookUC:           uc_book.NewBookUC(repo),
	}
}
