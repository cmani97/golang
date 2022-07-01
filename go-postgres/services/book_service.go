package services

import (
	"context"
	"fmt"

	"github.com/slice-mani/gin-postgres-project/models"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBook(*models.Book) error
	GetAllBooks() ([]*models.Book, error)
	GetBook(*string) (*models.Book, error)
}

type BookServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewBookService(db *gorm.DB, ctx context.Context) BookService {
	return &BookServiceImpl{
		db:  db,
		ctx: ctx,
	}
}

func (bookservice *BookServiceImpl) CreateBook(book *models.Book) error {
	fmt.Println("CreateBook service called")
	result := bookservice.db.Create(&book)
	fmt.Println("resutl= ", result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (bookservice *BookServiceImpl) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	err := bookservice.db.Find(&books)
	if err != nil {
		return nil, err.Error
	}
	return books, nil
}

func (bookservice *BookServiceImpl) GetBook(title *string) (*models.Book, error) {
	var book *models.Book
	if err := bookservice.db.Where("id = ?", title).First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
