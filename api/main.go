package main

import (
	"os"

	"github.com/kkk-petrov/gobooks-api/controllers"
	"github.com/kkk-petrov/gobooks-api/database"
	"github.com/kkk-petrov/gobooks-api/routes"
	"github.com/kkk-petrov/gobooks-api/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	PORT    string
	JWT_KEY string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env")
	}

	PORT = os.Getenv("PORT")
	JWT_KEY = os.Getenv("JWT_KEY")
}

func main() {
	e := echo.New()

	booksService := services.NewBooksService(database.DB)
	usersService := services.NewUsersService(database.DB)
	authService := services.NewAuthService(&usersService, []byte(JWT_KEY))

	booksController := controllers.NewBooksController(&booksService)
	usersController := controllers.NewUsersController(&usersService)
	authController := controllers.NewAuthController(&authService, &usersController)

	booksGroup := e.Group("/books")
	usersGroup := e.Group("/users")
	authGroup := e.Group("/auth")

	routes.SetupBookRoutes(booksGroup, &booksController)
	routes.SetupUserRoutes(usersGroup, &usersController)
	routes.SetupAuthRoutes(authGroup, &authController)

	e.Logger.Fatal(e.Start(":" + PORT))
}
