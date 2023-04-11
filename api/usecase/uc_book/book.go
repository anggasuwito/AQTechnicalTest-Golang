package uc_book

import (
	"AQTechincalTest-Golang/api/models"
	"AQTechincalTest-Golang/api/repository"
	"AQTechincalTest-Golang/helper/pagination"
)

type BookUC struct {
	Repo repository.Repository
}

type BookUCInterface interface {
	CreateUC(req models.BookDetails) (err error)
	GetAllUC(query pagination.Request) (res []models.BookDetails, meta pagination.Response, err error)
	GetByIDUC(id string) (res models.BookDetails, err error)
	UpdateUC(req models.Book) (err error)
	DeleteUC(req models.BookListID) (err error)
}

func NewBookUC(repo repository.Repository) BookUCInterface {
	return &BookUC{Repo: repo}
}
