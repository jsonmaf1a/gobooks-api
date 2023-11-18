package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kkk-petrov/gobooks-api/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user models.User) (string, error)
	Login(username, password string) (string, error)
	GenerateToken(data userData) (string, error)
}

type userData struct {
	name  string
	email string
	id    uint
}

type authService struct {
	usersService UsersService
	key          []byte
}

func NewAuthService(s UsersService, key []byte) authService {
	return authService{
		usersService: s,
		key:          key,
	}
}

func (s *authService) Register(user models.User) (string, error) {
	err := s.usersService.CreateUser(user)
	if err != nil {
		return "", err
	}

	data := userData{
		id:    user.ID,
		name:  user.Name,
		email: user.Email,
	}

	return s.GenerateToken(data)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.usersService.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	data := userData{
		id:    user.ID,
		name:  user.Name,
		email: user.Email,
	}

	return s.GenerateToken(data)
}

func (s *authService) GenerateToken(data userData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":       data,
		"expires_at": "14",
	})

	return token.SignedString(s.key)
}

func (s *authService) VerifyToken(tokenString string) (*userData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userData := claims["user"].(userData)
		return &userData, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
