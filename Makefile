# Makefile - для удобной работы с проектом

.PHONY: all run build clean swagger test

# Переменные
APP_NAME=AuthServices
DB_NAME=local_db

# Основная цель
all: clean swagger build

# Запуск приложения
run:
	go run main.go

# Сборка приложения
build:
	go build -o $(APP_NAME) main.go

# Очистка
clean:
	go clean
	rm -f $(APP_NAME)

# Генерация Swagger документации
swagger:
	swag init -g main.go

# Запуск тестов
test:
	go test -v ./...

# Миграция базы данных
db-migrate:
	go run main.go migrate

# Создание базы данных PostgreSQL (требует установленного psql)
db-create:
	psql -U postgres -c "CREATE DATABASE $(DB_NAME);"

# Удаление базы данных
db-drop:
	psql -U postgres -c "DROP DATABASE IF EXISTS $(DB_NAME);"

# Сброс и повторное создание базы данных
db-reset: db-drop db-create db-migrate