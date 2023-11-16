package controllers

import (
	"books/models"
	"books/services"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BooksController interface {
	GetAllBooks(ctx echo.Context) error
	GetBookById(ctx echo.Context) error
	CreateBook(ctx echo.Context) error
	UpdateBook(ctx echo.Context) error
	DeleteBook(ctx echo.Context) error
}

type booksController struct {
	booksService services.BooksService
}

func NewBooksController(s services.BooksService) booksController {
	return booksController{
		booksService: s,
	}
}

func (c booksController) GetAllBooks(ctx echo.Context) error {
	books, err := c.booksService.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, books)
}

func (c booksController) GetBookById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	book, err := c.booksService.FindById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, book)
}

func (c booksController) CreateBook(ctx echo.Context) error {
	var dto models.CreateBookDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	newBook := models.Book{
		Title:       dto.Title,
		Author:      dto.Author,
		Description: dto.Description,
		PublishedAt: dto.PublishedAt,
	}

	err := c.booksService.Create(newBook)
	if err != nil {
		log.Panic(err)
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "Book created successfully"})
}

func (c booksController) UpdateBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "book not found"})
	}

	var dto models.UpdateBookDto
	if err := ctx.Bind(&dto); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	err = c.booksService.Update(id, dto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Book updated successfully"})
}

func (c booksController) DeleteBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "book not found"})
	}

	err = c.booksService.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}
