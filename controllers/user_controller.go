
// controllers/user_controller.go - обработчики HTTP запросов для пользователей с Swagger аннотациями
package controllers

import (
	"net/http"

	"AuthApplications/services"
    "AuthApplications/dto"
    
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

// UserController интерфейс контроллера пользователей
type UserController interface {
	GetProfile(c *gin.Context)
	GetAllUsers(c *gin.Context)
	GetByID(c *gin.Context)
    PatchUser(c *gin.Context)
    DeleteUser(c *gin.Context)
}

// userController реализация UserController
type userController struct {
	userService services.UserService
}

// NewUserController создает новый контроллер пользователей
func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary Получение профиля пользователя
// @Description Возвращает информацию о профиле текущего аутентифицированного пользователя
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse "Профиль пользователя"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users/profile [get]
func (ctrl *userController) GetProfile(c *gin.Context) {
	// Получить ID пользователя из контекста (устанавливается middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	// Получить профиль пользователя
	profile, err := ctrl.userService.GetUserProfile(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetAllUsers godoc
// @Summary Получение всех пользователей
// @Description Возвращает список всех пользователей
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.UserResponse "Список пользователей"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users/all [get]
func (ctrl *userController) GetAllUsers(c *gin.Context) {
    // Получить всех пользователей через UserService
    users, err := ctrl.userService.GetAllUser()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, users)
}

// GetByID godoc
// @Summary Получение пользователя по ID
// @Description Возвращает информацию о пользователе по указанному ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse "Пользователь найден"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users/{id} [get]
func (ctrl *userController) GetByID(c *gin.Context) {
    // Получить ID из параметров маршрута
    idParam := c.Param("id")
    userID, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
        return
    }

    // Получить пользователя через UserService
    user, err := ctrl.userService.GetByID(userID)
    if err != nil {
        if err.Error() == "record not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, user)
}

// PatchUserRequsest godoc
// @Summary Полное обновление пользователя
// @Description Полностью обновляет данные пользователя по указанному ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Param user body dto.PatchUserRequsest true "Данные для обновления"
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse "Пользователь обновлен"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users/{id} [patch]
func (ctrl *userController) PatchUser(c *gin.Context) {
    idParam := c.Param("id")
    userID, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
        return
    }

    var request dto.PatchUserRequsest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := ctrl.userService.PatchUser(userID, request)
    if err != nil {
        if err.Error() == "record not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Удаление пользователя
// @Description Удаляет пользователя по указанному ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "ID пользователя"
// @Security BearerAuth
// @Success 204 "Пользователь успешно удален"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/users/{id} [delete]
func (ctrl *userController) DeleteUser(c *gin.Context) {
    idParam := c.Param("id")
    userID, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
        return
    }

    // Удаляем пользователя через сервис
    err = ctrl.userService.DeleteUser(userID)
    if err != nil {
        if err.Error() == "record not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Пользователь удален",
        "user_id": userID,
    })    
}
