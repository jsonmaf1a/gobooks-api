package services

import (
	"books/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BooksService interface {
	Create(book models.Book) error
	FindAll() ([]models.Book, error)
	FindById(id int) (models.Book, error)
	Update(id int, data models.UpdateBookDto) error
	Delete(id int) error
}

type booksService struct {
	db *gorm.DB
}

func NewBooksService(db *gorm.DB) booksService {
	return booksService{
		db: db,
	}
}

func (s booksService) Create(book models.Book) error {
	err := s.db.Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (s booksService) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := s.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s booksService) FindById(id int) (models.Book, error) {
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

func (s booksService) Update(id int, data models.UpdateBookDto) error {
	err := s.db.Model(&models.Book{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s booksService) Delete(id int) error {
	err := s.db.Delete(&models.Book{}, "id = ?", id).Error
	return err
}
