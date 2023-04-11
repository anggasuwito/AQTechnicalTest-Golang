package r_book

import "AQTechincalTest-Golang/api/models"

func (r BookRepo) CountAllRepo() (res int, err error) {
	err = r.DB.Debug().Model(&models.Book{}).Where("deleted_at IS NULL").Count(&res).Error
	return res, err
}
