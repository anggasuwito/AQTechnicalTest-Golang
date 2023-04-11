package r_book

import "AQTechincalTest-Golang/api/models"

func (r BookRepo) GetAllRepo(offset int, limit int) (res []models.Book, err error) {
	err = r.DB.Debug().Where("deleted_at IS NULL").Order("created_at DESC").Limit(limit).Offset(offset).Find(&res).Error
	return res, err
}
