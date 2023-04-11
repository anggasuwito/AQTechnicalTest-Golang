package r_book

import "AQTechincalTest-Golang/api/models"

func (r BookRepo) GetKeywordsRepo(bookID int) (res []models.BookKeyword, err error) {
	err = r.DB.Debug().Where("deleted_at IS NULL AND book_id = ?", bookID).Order("created_at DESC").Find(&res).Error
	return res, err
}
