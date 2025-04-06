package models

import (
	"time"
    "github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	AuthorID    uuid.UUID `gorm:"not null;foreignKey:UserID" json:"user_id"`
	Description string    `json:"description"`
	ISBN        string    `gorm:"unique" json:"isbn"`
	PublishYear int       `json:"publish_year"`
	CoverURL    string    `json:"cover_url"`
	FileURL     string    `json:"file_url"`
	Genre       string    `json:"genre"`
	Language    string    `json:"language"`
	PageCount   int       `json:"page_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}