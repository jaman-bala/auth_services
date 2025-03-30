// main.go - точка входа в приложение
package main

import (
	"log"

	_ "AuthApplications/docs"
	"AuthApplications/config"
	"AuthApplications/routes"
	
)

// @title Auth Services API
func main() {
	// Инициализация конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Инициализация базы данных
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Настройка и запуск роутера
	r := routes.SetupRouter(db, cfg)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}