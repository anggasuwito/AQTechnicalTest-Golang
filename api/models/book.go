package models

import "time"

func (Book) TableName() string {
	return "book"
}

type Book struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	Harga     float64   `json:"harga"`
	Stok      int       `json:"stok"`
	Penerbit  string    `json:"penerbit"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (BookCategory) TableName() string {
	return "book_category"
}

type BookCategory struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	BookID    int       `json:"book_id"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (BookKeyword) TableName() string {
	return "book_keyword"
}

type BookKeyword struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	BookID    int       `json:"book_id"`
	Keyword   string    `json:"keyword"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type BookListID struct {
	ID []int `json:"id"`
}

type BookDetails struct {
	Book
	BookCategory []BookCategory `json:"book_category"`
	BookKeyword  []BookKeyword  `json:"book_keyword"`
}
