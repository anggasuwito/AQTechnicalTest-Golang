package r_book

import (
	"AQTechincalTest-Golang/api/models"
	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

type BookRepoInterface interface {
	TXCreateRepo(req models.BookDetails) (err error)
	TXDeleteRepo(id models.BookListID) (err error)
	UpdateRepo(req models.Book) (err error)
	CountAllRepo() (res int, err error)
	GetAllRepo(offset int, limit int) (res []models.Book, err error)
	GetByIDRepo(bookID int) (res models.Book, err error)
	GetCategoriesRepo(bookID int) (res []models.BookCategory, err error)
	GetKeywordsRepo(bookID int) (res []models.BookKeyword, err error)
}

func NewBookRepo(db *gorm.DB) BookRepoInterface {
	return &BookRepo{DB: db}
}
