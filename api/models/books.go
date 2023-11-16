package models

import "time"

type Book struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	PublishedAt int       `json:"published_at"`
}

type CreateBookDto struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublishedAt int    `json:"published_at"`
}

type UpdateBookDto struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PublishedAt int    `json:"published_at"`
}
