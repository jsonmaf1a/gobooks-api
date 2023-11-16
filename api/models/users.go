package models

import "time"

type User struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	ID           uint      `gorm:"primaryKey;autoIncrement"`
}

type CreateUserDto struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type UpdateUserDto struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}
