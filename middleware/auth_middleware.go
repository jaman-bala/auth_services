// middleware/auth_middleware.go - промежуточное ПО для аутентификации
package middleware

import (
	"net/http"

	"AuthApplications/services"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware для проверки JWT токена
func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получить токен из заголовка Authorization
		tokenString, err := c.Cookie("auth_token")
		if err != nil || tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
			c.Abort()
			return
		}

		// Валидация токена
		_, claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен: " + err.Error()})
			c.Abort()
			return
		}

		// Устанавливаем данные пользователя в контексте
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleMiddleware middleware для проверки роли пользователя
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получить роль из контекста (установлена AuthMiddleware)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			c.Abort()
			return
		}

		// Проверить, имеет ли пользователь необходимую роль
		userRole := role.(string)
		hasRole := false
		for _, r := range roles {
			if r == userRole {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "У вас нет прав для доступа к этому ресурсу"})
			c.Abort()
			return
		}

		c.Next()
	}
}