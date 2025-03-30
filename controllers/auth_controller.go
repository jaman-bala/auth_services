// controllers/auth_controller.go - обработчики HTTP запросов для аутентификации с Swagger аннотациями
package controllers

import (
	"net/http"

	"AuthApplications/dto"
	"AuthApplications/services"
	"AuthApplications/config"
	"github.com/gin-gonic/gin"
)

// AuthController интерфейс контроллера аутентификации
type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

// authController реализация AuthController
type authController struct {
	authService services.AuthService
}

// NewAuthController создает новый контроллер аутентификации
func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "Данные пользователя"
// @Success 201 {object} map[string]interface{} "Пользователь успешно зарегистрирован"
// @Failure 400 {object} map[string]string "Ошибка валидации или пользователь уже существует"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/auth/register [post]
func (ctrl *authController) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.authService.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"user_id": user.ID,
	})
}

// Login godoc
// @Summary Вход в систему
// @Description Аутентифицирует пользователя и возвращает JWT токен
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Учетные данные"
// @Success 200 {object} dto.AuthResponse "Успешный вход в систему"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 401 {object} map[string]string "Неверные учетные данные"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/auth/login [post]
func (ctrl *authController) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.authService.Login(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	config, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки конфигурации"})
		return
	}

	c.SetCookie("auth_token", token, config.CookieLifetime, "/", config.CookieDomain, true, true)

	response := dto.AuthResponse{
		Token:   token,
		Message: "Успешный вход в систему",
	}

	c.JSON(http.StatusOK, response)
}

// Logout godoc
// @Summary Выход из системы
// @Description Удаляет токен и завершает сессию пользователя
// @Tags auth
// @Security BearerAuth
// @Success 200 {object} map[string]string "Успешный выход из системы"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/auth/logout [post]
func (ctrl *authController) Logout(c *gin.Context) {
    err := ctrl.authService.Logout()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выходе из системы"})
        return
    }

    // Очистка токена в cookies
    c.SetCookie("auth_token", "", -1, "/", "", true, true)

    c.JSON(http.StatusOK, gin.H{
        "message": "Успешный выход из системы",
    })
}