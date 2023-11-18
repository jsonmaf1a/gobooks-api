package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kkk-petrov/gobooks-api/models"
	"github.com/kkk-petrov/gobooks-api/services"

	"github.com/labstack/echo/v4"
)

type BooksController interface {
	FindAll(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type booksController struct {
	booksService services.BooksService
}

func NewBooksController(s services.BooksService) booksController {
	return booksController{
		booksService: s,
	}
}

func (c *booksController) FindAll(ctx echo.Context) error {
	books, err := c.booksService.FindAllBooks()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, books)
}

func (c *booksController) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	book, err := c.booksService.FindBookById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, book)
}

func (c *booksController) Create(ctx echo.Context) error {
	var dto models.CreateBookDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	newBook := models.Book{
		Title:       dto.Title,
		Author:      dto.Author,
		Description: dto.Description,
		PublishedAt: dto.PublishedAt,
	}

	err := c.booksService.CreateBook(newBook)
	if err != nil {
		log.Panic(err)
	}

	return ctx.JSON(http.StatusCreated, Response{"Book created successfully"})
}

func (c *booksController) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"book not found"})
	}

	var dto models.UpdateBookDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, Response{"Invalid request payload"})
	}

	err = c.booksService.UpdateBook(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}
	return ctx.JSON(http.StatusOK, Response{"Book updated successfully"})
}

func (c *booksController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{"book not found"})
	}

	err = c.booksService.DeleteBook(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, Response{err.Error()})
	}

	return ctx.JSON(http.StatusOK, Response{"Book deleted successfully"})
}
