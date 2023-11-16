package database

import (
	"books/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_SSLMODE  string
	ConnStr     string
	DB          *gorm.DB
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env")
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_SSLMODE = os.Getenv("DB_SSLMODE")

	ConnStr = fmt.Sprintf("host=%v user=%v database=%v password=%v sslmode=%v", DB_HOST, DB_USER, DB_NAME, DB_PASSWORD, DB_SSLMODE)

	var err error
	DB, err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	DB.AutoMigrate(&models.Book{}, &models.User{})
}
