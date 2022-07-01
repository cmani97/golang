package models

type Book struct {
	ID       uint   `json:"book_id" gorm:"primary_key"`
	Title    string `json:"title"`
	Desc     string `json:"description"`
	AuthorID uint   `json:"author_id" `
	Author   Author `gorm:"foreignKey:AuthorID;references:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
