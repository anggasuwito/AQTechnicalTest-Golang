package uc_book

import (
	"AQTechincalTest-Golang/api/models"
)

func (uc BookUC) UpdateUC(req models.Book) (err error) {
	return uc.Repo.BookRepo.UpdateRepo(req)
}
