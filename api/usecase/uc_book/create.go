package uc_book

import (
	"AQTechincalTest-Golang/api/models"
)

func (uc BookUC) CreateUC(req models.BookDetails) (err error) {
	return uc.Repo.BookRepo.TXCreateRepo(req)
}
