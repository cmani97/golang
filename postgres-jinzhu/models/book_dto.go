package models

type BookDto struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Desc   string `json:"description" binding:"required"`
}
