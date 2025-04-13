// config/database.go - настройка базы данных
package config

import (
	"fmt"

	"AuthApplications/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB инициализирует подключение к базе данных
func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция таблиц
 	err = db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.AuthorBook{},
		)
	if err != nil {
		return nil, err
	}

	return db, nil
}