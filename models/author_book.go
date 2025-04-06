package models

import (
	"time"
	"github.com/google/uuid"
)


type AuthorBook struct {
	ID        	uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AuthorID    uuid.UUID `gorm:"not null;foreignKey:UserID" json:"user_id"`
	BookID    	uuid.UUID `gorm:"not null;foreignKey:BookID" json:"book_id"`
	IsFavorite 	bool 	  `gorm:"default:false" json:"is_favorite"`
	ReadingState string `gorm:"default:'not_started'" json:"reading_state"`
	ReadProgress int `gorm:"default:0" json:"reading_progress"`
	LastReadPage int `json:"last_read_page"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}