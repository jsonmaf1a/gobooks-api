package services

import (
	"errors"
	"fmt"

	"github.com/kkk-petrov/gobooks-api/models"

	"gorm.io/gorm"
)

type BooksService interface {
	CreateBook(book models.Book) error
	FindAllBooks() ([]models.Book, error)
	FindBookById(id int) (models.Book, error)
	UpdateBook(id int, data models.UpdateBookDto) error
	DeleteBook(id int) error
}

type booksService struct {
	db *gorm.DB
}

func NewBooksService(db *gorm.DB) booksService {
	return booksService{
		db: db,
	}
}

func (s *booksService) CreateBook(book models.Book) error {
	err := s.db.Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *booksService) FindAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := s.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *booksService) FindBookById(id int) (models.Book, error) {
	var book models.Book

	err := s.db.First(&book, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("book with ID %d not found", id)
		}
		return models.Book{}, err
	}
	return book, nil
}

func (s *booksService) UpdateBook(id int, data models.UpdateBookDto) error {
	err := s.db.Model(&models.Book{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *booksService) DeleteBook(id int) error {
	err := s.db.Delete(&models.Book{}, "id = ?", id).Error
	return err
}
