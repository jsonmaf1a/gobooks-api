package main

import (
	"books/controllers"
	"books/database"
	"books/routes"
	"books/services"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env")
	}

	PORT = os.Getenv("PORT")
}

func main() {
	e := echo.New()

	booksService := services.NewBooksService(database.DB)
	booksController := controllers.NewBooksController(booksService)
	booksGroup := e.Group("/books")

	routes.SetupBookRoutes(booksGroup, booksController)

	e.Logger.Fatal(e.Start(":" + PORT))
}
