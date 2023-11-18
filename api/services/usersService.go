package services

import (
	"github.com/kkk-petrov/gobooks-api/models"

	"gorm.io/gorm"
)

type UsersService interface {
	CreateUser(user models.User) error
	FindAllUsers() ([]models.User, error)
	FindUserById(id int) (models.User, error)
	FindUserByEmail(email string) (models.User, error)
	UpdateUser(id int, data models.UpdateUserDto) error
	DeleteUser(id int) error
}

type usersService struct {
	db *gorm.DB
}

func NewUsersService(db *gorm.DB) usersService {
	return usersService{
		db: db,
	}
}

func (s *usersService) CreateUser(user models.User) error {
	err := s.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *usersService) FindAllUsers() ([]models.User, error) {
	var users []models.User

	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *usersService) FindUserById(id int) (models.User, error) {
	var user models.User

	err := s.db.Find(&user, "id = ?", id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *usersService) FindUserByEmail(email string) (models.User, error) {
	var user models.User

	err := s.db.Find(&user, "email = ?", email).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *usersService) UpdateUser(id int, data models.UpdateUserDto) error {
	err := s.db.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *usersService) DeleteUser(id int) error {
	err := s.db.Delete(&models.User{}, "id = ?", id).Error
	return err
}
