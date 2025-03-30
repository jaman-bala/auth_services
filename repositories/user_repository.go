package repositories

import (
	"AuthApplications/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]models.User, error) 
	PatchUser(user *models.User) error
	DeleteByID(id uuid.UUID) error
}

// userRepository реализация UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository создает новый репозиторий пользователей
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create создает нового пользователя
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindAll возвращает всех пользователей
func (r *userRepository) FindAll() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error // GORM метод для выборки всех записей
    if err != nil {
        return nil, err
    }
    return users, nil
}

// FindByEmail находит пользователя по email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID находит пользователя по ID
func (r *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update обновляет пользователя в базе данных
func (r *userRepository) PatchUser(user *models.User) error {
    return r.db.Save(user).Error
}

// DeleteByID удаляет пользователя из базы данных по ID
func (r *userRepository) DeleteByID(id uuid.UUID) error {
    return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
