// dto/auth.go - структуры для передачи данных
package dto


// RegisterRequest представляет запрос на регистрацию
type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// LoginRequest представляет запрос на вход
type LoginRequest struct {
	Email string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"string"`
}

// AuthResponse представляет ответ после аутентификации
type AuthResponse struct {
	Token   string `json:"token"`
	Message string `json:"message,omitempty"`
}

// UserResponse представляет информацию о пользователе в ответе
type UserResponse struct {
	ID        string   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type PatchUserRequsest struct {
    Username  *string `json:"username,omitempty"`
    Email     *string `json:"email,omitempty"`
    FirstName *string `json:"first_name,omitempty"`
    LastName  *string `json:"last_name,omitempty"`
    Role      *string `json:"role,omitempty"`
}
