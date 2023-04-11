package uc_book

import (
	"AQTechincalTest-Golang/api/models"
)

func (uc BookUC) DeleteUC(req models.BookListID) (err error) {
	return uc.Repo.BookRepo.TXDeleteRepo(req)
}
