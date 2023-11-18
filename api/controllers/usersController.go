package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kkk-petrov/gobooks-api/models"
	"github.com/kkk-petrov/gobooks-api/services"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
}

type UsersController interface {
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type usersController struct {
	usersService services.UsersService
}

func NewUsersController(s services.UsersService) usersController {
	return usersController{
		usersService: s,
	}
}

func (c *usersController) FindAll(ctx echo.Context) error {
	users, err := c.usersService.FindAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, users)
}

func (c *usersController) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	user, err := c.usersService.FindUserById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, user)
}

func (c *usersController) Create(ctx echo.Context) error {
	var dto models.CreateUserDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
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

	err = c.usersService.CreateUser(newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, Response{"User created successfully"})
}

func (c *usersController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}

	var dto models.UpdateUserDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	err = c.usersService.UpdateUser(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}
	return ctx.JSON(http.StatusOK, Response{"User updated successfully"})
}

func (c *usersController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}

	err = c.usersService.DeleteUser(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, Response{"User deleted successfully"})
}
