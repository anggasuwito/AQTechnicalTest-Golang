package r_book

import (
	"AQTechincalTest-Golang/api/models"
	"time"
)

func (r BookRepo) UpdateRepo(req models.Book) (err error) {
	return r.DB.Debug().Model(&models.Book{}).Where("deleted_at IS NULL and id = ?", req.ID).Updates(map[string]interface{}{
		"judul":      req.Judul,
		"deskripsi":  req.Deskripsi,
		"harga":      req.Harga,
		"stok":       req.Stok,
		"penerbit":   req.Penerbit,
		"updated_at": time.Now().Format(time.RFC3339),
	}).Error
}
