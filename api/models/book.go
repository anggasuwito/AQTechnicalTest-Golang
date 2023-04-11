package models

func (Book) TableName() string {
	return "book"
}

type Book struct {
	ID        int     `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float64 `json:"harga"`
	Stok      int     `json:"stok"`
	Penerbit  string  `json:"penerbit"`
	CreatedAt string  `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	UpdatedAt string  `json:"updated_at" gorm:"type:datetime"`
	DeletedAt string  `json:"deleted_at" gorm:"type:datetime"`
}

func (BookCategory) TableName() string {
	return "book_category"
}

type BookCategory struct {
	ID        int    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	BookID    int    `json:"book_id"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime"`
	DeletedAt string `json:"deleted_at" gorm:"type:datetime"`
}

func (BookKeyword) TableName() string {
	return "book_keyword"
}

type BookKeyword struct {
	ID        int    `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT_NULL"`
	BookID    int    `json:"book_id"`
	Keyword   string `json:"keyword"`
	CreatedAt string `json:"created_at" gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	UpdatedAt string `json:"updated_at" gorm:"type:datetime"`
	DeletedAt string `json:"deleted_at" gorm:"type:datetime"`
}

type BookListID struct {
	ID []int `json:"id"`
}

type BookDetails struct {
	Book
	BookCategory []BookCategory `json:"book_category"`
	BookKeyword  []BookKeyword  `json:"book_keyword"`
}
