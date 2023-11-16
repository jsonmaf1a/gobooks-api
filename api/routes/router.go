package routes

import (
	"books/controllers"

	"github.com/labstack/echo/v4"
)

func SetupBookRoutes(booksGroup *echo.Group, c controllers.BooksController) {
	booksGroup.GET("", c.GetAllBooks)
	booksGroup.GET("/:id", c.GetBookById)
	booksGroup.POST("", c.CreateBook)
	booksGroup.PUT("/:id", c.UpdateBook)
	booksGroup.DELETE("/:id", c.DeleteBook)
}

func SetupUserRoutes(booksGroup *echo.Group, c controllers.UsersController) {
	booksGroup.POST("", c.Login)
	booksGroup.POST("", c.Register)
}
