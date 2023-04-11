package uc_book

import (
	"AQTechincalTest-Golang/api/models"
	"strconv"
)

func (uc BookUC) GetByIDUC(id string) (res models.BookDetails, err error) {
	numID, _ := strconv.Atoi(id)
	book, err := uc.Repo.BookRepo.GetByIDRepo(numID)
	if err != nil {
		return res, err
	}

	categories, err := uc.Repo.BookRepo.GetCategoriesRepo(numID)
	if err != nil {
		return res, err
	}

	keywords, err := uc.Repo.BookRepo.GetKeywordsRepo(numID)
	if err != nil {
		return res, err
	}

	res.Book = book
	res.BookCategory = categories
	res.BookKeyword = keywords
	return res, err
}
