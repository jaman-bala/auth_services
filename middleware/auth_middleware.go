// middleware/auth_middleware.go - промежуточное ПО для аутентификации
package middleware

import (
	"net/http"

	"AuthApplications/services"
	"github.com/gin-gonic/gin"
)

const (
    AuthorizationHeaderKey = "Authorization"
    BearerSchema           = "Bearer "
    AccessTokenCookieName  = "access_token"
)


// AuthMiddleware middleware для проверки JWT токена
func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
        var tokenString string

        // Попробовать получить токен из cookie
        cookieToken, err := c.Cookie(AccessTokenCookieName)
        if err == nil && cookieToken != "" {
            tokenString = cookieToken
        } else {
            // Если в cookie нет — пробуем из заголовка Authorization
            authHeader := c.GetHeader(AuthorizationHeaderKey)
            if len(authHeader) > len(BearerSchema) && authHeader[:len(BearerSchema)] == BearerSchema {
                tokenString = authHeader[len(BearerSchema):]
            }
        }

		// Если токен всё ещё пуст — возвращаем ошибку
		if tokenString == "" {
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

		// Устанавливаем данные пользователя в контекст
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