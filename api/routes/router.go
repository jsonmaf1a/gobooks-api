package routes

import (
	"github.com/kkk-petrov/gobooks-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupBookRoutes(booksGroup *echo.Group, c controllers.BooksController) {
	booksGroup.GET("", c.FindAll)
	booksGroup.GET("/:id", c.FindById)
	booksGroup.POST("", c.Create)
	booksGroup.PUT("/:id", c.Update)
	booksGroup.DELETE("/:id", c.Delete)
}

func SetupUserRoutes(usersGroup *echo.Group, c controllers.UsersController) {
	usersGroup.GET("", c.FindAll)
	usersGroup.GET("/:id", c.FindById)
	usersGroup.POST("", c.Create)
	usersGroup.PUT("/:id", c.Update)
	usersGroup.DELETE("/:id", c.Delete)
}

func SetupAuthRoutes(authGroup *echo.Group, c controllers.AuthController) {
	authGroup.POST("/register", c.Register)
	authGroup.POST("/login", c.Login)
}
