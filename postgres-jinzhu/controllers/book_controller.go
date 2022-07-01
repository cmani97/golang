package controllers

import (
	"fmt"
	"net/http"

	models "github.com/slice-mani/postgres-jinzhu/models"
	"github.com/slice-mani/postgres-jinzhu/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	BookService services.BookService
}

func NewBookController(bookservice services.BookService) BookController {
	return BookController{
		BookService: bookservice,
	}
}

func (bookcontroller *BookController) FindBooks(ctx *gin.Context) {
	books, err := bookcontroller.BookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func (bookcontroller *BookController) CreateBook(ctx *gin.Context) {
	fmt.Println("CreateBook controller called")
	var input models.BookDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	book := models.Book{Title: input.Title, Desc: input.Desc, AuthorID: 1}
	fmt.Println("book= ", book)
	err := bookcontroller.BookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "err.Error()"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (bookcontroller *BookController) FindBook(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := bookcontroller.BookService.GetBook(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func (bookcontroller *BookController) UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.BookDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (bookcontroller *BookController) DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
