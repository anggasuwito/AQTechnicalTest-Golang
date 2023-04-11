package uc_book

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/helper/pagination"
)

func (uc BookUC) GetAllUC(query pagination.Request) (res []models.BookDetails, meta pagination.Response, err error) {
	offset, limit := pagination.GetPagination(query)
	books, err := uc.Repo.BookRepo.GetAllRepo(offset, limit)
	if err != nil {
		return res, meta, err
	}
	for _, book := range books {
		categories, err := uc.Repo.BookRepo.GetCategoriesRepo(book.ID)
		if err != nil {
			return res, meta, err
		}

		keywords, err := uc.Repo.BookRepo.GetKeywordsRepo(book.ID)
		if err != nil {
			return res, meta, err
		}

		res = append(res, models.BookDetails{
			Book:         book,
			BookCategory: categories,
			BookKeyword:  keywords,
		})
	}

	total, err := uc.Repo.BookRepo.CountAllRepo()
	if err != nil {
		return res, meta, err
	}

	meta = pagination.GetPaginationResponse(query.Page, limit, total)
	return res, meta, err
}
