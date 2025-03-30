// models/user.go - модель пользователя
package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User представляет модель пользователя в базе данных
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // Не отображать пароль в JSON
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `gorm:"default:user" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeSave хеширует пароль перед сохранением пользователя
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Хешируем пароль только если он был изменен
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword проверяет, совпадает ли введенный пароль с хешированным
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}