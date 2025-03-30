// services/auth_service.go - бизнес-логика аутентификации
package services

import (
	"errors"
	"time"

	"AuthApplications/config"
	"AuthApplications/dto"
	"AuthApplications/models"
	"AuthApplications/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthService интерфейс сервиса аутентификации
type AuthService interface {
	Register(req dto.RegisterRequest) (*models.User, error)
	Login(req dto.LoginRequest) (string, error)
	Logout() error
	ValidateToken(tokenString string) (*jwt.Token, *JWTClaim, error)
}

// JWTClaim представляет структуру JWT токена
type JWTClaim struct {
	UserID   uuid.UUID    `json:"user_id"`
	Email string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// authService реализация AuthService
type authService struct {
	userRepo  repositories.UserRepository
	jwtSecret string
}

// NewAuthService создает новый сервис аутентификации
func NewAuthService(userRepo repositories.UserRepository, cfg *config.Config) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: cfg.JWTSecret,
	}
}

// Register регистрирует нового пользователя
func (s *authService) Register(req dto.RegisterRequest) (*models.User, error) {
	// Проверка, существует ли пользователь с таким email
	_, err := s.userRepo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("пользователь с таким email уже существует")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Создание нового пользователя
	newUser := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      "user", // По умолчанию обычный пользователь
	}

	if err := s.userRepo.Create(newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

// Login аутентифицирует пользователя и выдает JWT токен
func (s *authService) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("неверное имя пользователя или пароль")
		}
		return "", err
	}

	// Проверка пароля
	if err := user.CheckPassword(req.Password); err != nil {
		return "", errors.New("неверное имя пользователя или пароль")
	}

	// Создание JWT токена
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID:   user.ID,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Logout реализация метода для выхода из системы
func (s *authService) Logout() error {
    // Здесь можно настроить дополнительную логику, например очистку данных пользователя
    // В данном случае достаточно удалить токен из cookies (будет обработано в контроллере)
    return nil
}

// ValidateToken проверяет и валидирует JWT токен
func (s *authService) ValidateToken(tokenString string) (*jwt.Token, *JWTClaim, error) {
	claims := &JWTClaim{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.jwtSecret), nil
		},
	)

	if err != nil {
		return nil, nil, errors.New("ошибка при разборе токена: " + err.Error())
	}

	if !token.Valid {
		return nil, nil, errors.New("недействительный токен")
	}

	return token, claims, nil
}