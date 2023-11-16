package services

import (
	"books/models"

	"gorm.io/gorm"
)

type UsersService interface {
	Create(book models.User) error
	FindAll() ([]models.User, error)
	FindById(id int) (models.User, error)
	Update(id int, data models.UpdateUserDto) error
	Delete(id int) error
}

type usersService struct {
	db *gorm.DB
}

func NewUsersService(db *gorm.DB) booksService {
	return booksService{
		db: db,
	}
}

func (s usersService) Create() error {
	return nil
}

func (s usersService) FindAll() ([]models.User, error) {
	var users []models.User

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s usersService) FindById(id int) (models.User, error) {
	var user models.User

	err := s.db.Find(&user, "id = ?", id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s usersService) Update(id int, data models.UpdateUserDto) error {
	err := s.db.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s usersService) Delete(id int) error {
	err := s.db.Delete(&models.User{}, "id = ?", id).Error
	return err
}
