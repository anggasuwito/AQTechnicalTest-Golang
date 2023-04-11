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

	err = tx.Debug().Select(
		"BookID",
		"Category",
	).Create(&req.BookCategory).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Select(
		"BookID",
		"Keyword",
	).Create(&req.BookKeyword).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
