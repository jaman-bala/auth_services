# Auth API Project

Проект API для аутентификации и авторизации пользователей на основе Go, Gin и JWT.

## Особенности

- Регистрация пользователей
- Аутентификация с JWT
- Контроль доступа на основе ролей
- Swagger документация API
- Многоуровневая архитектура

## Требования

- Go 1.20 или выше
- PostgreSQL
- Make (опционально)

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/username/auth-project.git
cd auth-project
```

### 2. Установка зависимостей

```bash
go mod download
```

### 3. Установка Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 4. Создание файла .env

Создайте файл `.env` в корневой директории проекта и добавьте в него следующие параметры:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=authdb
JWT_SECRET=your_super_secret_key_here
SERVER_PORT=8080
```

### 5. Создание базы данных

```bash
make db-create
```
или вручную:
```bash
psql -U postgres -c "CREATE DATABASE authdb;"
```

### 6. Генерация Swagger документации

```bash
make swagger
```
или вручную:
```bash
swag init -g main.go
```

### 7. Запуск приложения

```bash
make run
```
или вручную:
```bash
go run main.go
```

## Доступ к Swagger UI

После запуска приложения, Swagger UI доступен по адресу:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Публичные маршруты:

- **POST /api/auth/register** - Регистрация нового пользователя
- **POST /api/auth/login** - Вход в систему и получение JWT токена

### Защищенные маршруты (требуется JWT токен):

- **GET /api/users/profile** - Получение профиля текущего пользователя

### Маршруты администратора (требуется JWT токен с ролью admin):

- **GET /api/admin/** - Различные административные операции

## Структура проекта

```
auth-project/
├── main.go                 # Точка входа
├── .env                    # Файл с переменными окружения
├── Makefile                # Makefile для удобной работы
├── config/                 # Конфигурация приложения
├── controllers/            # Обработчики HTTP запросов
├── docs/                   # Swagger документация
├── dto/                    # Объекты передачи данных
├── middleware/             # Промежуточное ПО
├── models/                 # Модели данных
├── repositories/           # Слой доступа к данным
├── routes/                 # Маршруты
└── services/               # Бизнес-логика
```