package controllers

import (
	"net/http"

	"github.com/kkk-petrov/gobooks-api/models"
	"github.com/kkk-petrov/gobooks-api/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type authController struct {
	authService     services.AuthService
	usersController UsersController
}

func NewAuthController(s services.AuthService, c UsersController) authController {
	return authController{
		authService:     s,
		usersController: c,
	}
}

func (c *authController) Register(ctx echo.Context) error {
	var dto models.CreateUserDto
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	password := dto.PasswordHash

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"Error hashing password"})
	}

	newUser := models.User{
		Name:         dto.Name,
		Email:        dto.Email,
		PasswordHash: string(hashedPassword),
	}

	token, err := c.authService.Register(newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, Response{token})
}

func (c *authController) Login(ctx echo.Context) error {
	var loginDto struct {
		Email    string
		Password string
	}
	if err := ctx.Bind(&loginDto); err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	token, err := c.authService.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, Response{"Invalid email or password"})
	}

	return ctx.JSON(http.StatusOK, Response{token})
}
