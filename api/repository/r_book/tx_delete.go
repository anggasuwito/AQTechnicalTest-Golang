package r_book

import (
	"AQTechincalTest-Golang/api/models"
	"time"
)

func (r BookRepo) TXDeleteRepo(id models.BookListID) (err error) {
	tx := r.DB.Begin()
	err = tx.Debug().Model(&models.Book{}).Where("deleted_at IS NULL AND id IN (?)", id.ID).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Model(&models.BookCategory{}).Where("deleted_at IS NULL AND book_id IN (?)", id.ID).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Debug().Model(&models.BookKeyword{}).Where("deleted_at IS NULL AND book_id IN (?)", id.ID).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
