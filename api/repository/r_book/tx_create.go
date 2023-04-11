package r_book

import (
	"AQTechincalTest-Golang/api/models"
)

func (r BookRepo) TXCreateRepo(req models.BookDetails) (err error) {
	tx := r.DB.Begin()
	err = tx.Debug().Select(
		"Judul",
		"Deskripsi",
		"Harga",
		"Stok",
		"Penerbit",
	).Create(&req).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, category := range req.BookCategory {
		category.BookID = req.ID
		err = tx.Debug().Select(
			"BookID",
			"Category",
		).Create(&category).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, keyword := range req.BookKeyword {
		keyword.BookID = req.ID
		err = tx.Debug().Select(
			"BookID",
			"Keyword",
		).Create(&keyword).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
