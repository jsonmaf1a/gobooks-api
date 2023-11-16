package controllers

import (
	"books/services"

	"github.com/labstack/echo/v4"
)

type UsersController interface {
	GetAllUsers(ctx echo.Context) error
	GetUserById(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type usersController struct {
	usersService services.UsersService
}

func NewUsersController(s services.UsersService) usersController {
	return usersController{
		usersService: s,
	}
}

func (c usersController) Login(ctx echo.Context) error {
	return nil
}

func (c usersController) Register(ctx echo.Context) error {
	return nil
}
