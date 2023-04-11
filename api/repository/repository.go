package repository

import (
	"AQTechincalTest-Golang/api/repository/r_book"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	BookRepo r_book.BookRepoInterface
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		BookRepo: r_book.NewBookRepo(db),
	}
}
