package r_book

import "AQTechincalTest-Golang/api/models"

func (r BookRepo) GetByIDRepo(bookID int) (res models.Book, err error) {
	err = r.DB.Debug().Where("deleted_at IS NULL AND id = ?", bookID).Find(&res).Error
	return res, err
}
