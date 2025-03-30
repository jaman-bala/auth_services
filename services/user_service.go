// services/user_service.go - бизнес-логика для пользователей
package services

import (
	"AuthApplications/dto"
	"AuthApplications/repositories"

	"github.com/google/uuid"
)

// UserService интерфейс сервиса пользователей
type UserService interface {
	GetUserProfile(userID uuid.UUID) (*dto.UserResponse, error)
	GetAllUser() ([]*dto.UserResponse, error)
	GetByID(userID uuid.UUID) (*dto.UserResponse, error)
    PatchUser(userID uuid.UUID, req dto.PatchUserRequsest) (*dto.UserResponse, error) 
    DeleteUser(userID uuid.UUID) error
}

// userService реализация UserService
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService создает новый сервис пользователей
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// GetAllUser получает всех пользователей
func (s *userService) GetAllUser() ([]*dto.UserResponse, error) {
    users, err := s.userRepo.FindAll()
    if err != nil {
        return nil, err
    }

    // Преобразуем пользователей в массив DTO
    var userResponses []*dto.UserResponse
    for _, user := range users {
        userResponses = append(userResponses, &dto.UserResponse{
            ID:        user.ID.String(),
            Username:  user.Username,
            Email:     user.Email,
            FirstName: user.FirstName,
            LastName:  user.LastName,
            Role:      user.Role,
        })
    }

    return userResponses, nil
}

// GetUserProfile получает профиль пользователя
func (s *userService) GetUserProfile(userID uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}, nil
}

// GetByID находит пользователя по ID
func (s *userService) GetByID(userID uuid.UUID) (*dto.UserResponse, error) {
    user, err := s.userRepo.FindByID(userID) // Репозиторий должен реализовывать метод FindByID
    if err != nil {
        return nil, err
    }

    return &dto.UserResponse{
        ID:        user.ID.String(),
        Username:  user.Username,
        Email:     user.Email,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Role:      user.Role,
    }, nil
}

// UpdateUser обновляет данные пользователя
func (s *userService) PatchUser(userID uuid.UUID, req dto.PatchUserRequsest) (*dto.UserResponse, error) {
    // Найдем пользователя по ID
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return nil, err
    }

    // Обновляем только предоставленные поля
    if req.Username != nil {
        user.Username = *req.Username
    }
    if req.Email != nil {
        user.Email = *req.Email
    }
    if req.FirstName != nil {
        user.FirstName = *req.FirstName
    }
    if req.LastName != nil {
        user.LastName = *req.LastName
    }
    if req.Role != nil {
        user.Role = *req.Role
    }
    

    // Сохраним обновления
    if err := s.userRepo.PatchUser(user); err != nil {
        return nil, err
    }

    return &dto.UserResponse{
        ID:        user.ID.String(),
        Username:  user.Username,
        Email:     user.Email,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Role:      user.Role,
    }, nil
}

// DeleteUser удаляет пользователя по ID
func (s *userService) DeleteUser(userID uuid.UUID) error {
    // Проверим, существует ли пользователь
    _, err := s.userRepo.FindByID(userID)
    if err != nil {
        return err
    }

    // Удаляем пользователя
    return s.userRepo.DeleteByID(userID)
}
