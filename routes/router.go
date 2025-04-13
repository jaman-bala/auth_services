// routes/router.go - настройка маршрутов
package routes

import (
	"AuthApplications/config"
	"AuthApplications/controllers"
	"AuthApplications/middleware"
	"AuthApplications/repositories"
	"AuthApplications/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter настраивает и возвращает Gin router
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "GET /")
	})

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepo, cfg)
	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)

	// Инициализация контроллеров
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	bookController := controllers.NewBookController(bookService)

	// Публичные маршруты
	r.POST("/api/auth/register", authController.Register)
	r.POST("/api/auth/login", authController.Login)
	r.POST("/api/auth/logout", authController.Logout)

	// Группа защищенных маршрутов
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(authService))
	{
		// Маршруты пользователя
		protected.GET("/users/profile", userController.GetProfile)
		protected.GET("/users/all", userController.GetAllUsers)
		protected.GET("/users/:id", userController.GetByID)
		protected.PATCH("/users/:id", userController.PatchUser)
		protected.DELETE("/users/:id", userController.DeleteUser)
		
		// Маршруты книги
		protected.POST("/books", bookController.CreateBook)
        protected.GET("/books", bookController.GetAllBooks)
        protected.GET("/books/:id", bookController.GetByID)
		protected.GET("/books/genre/:genre", bookController.FindByGenre)
        // protected.DELETE("/books/:id", bookController.DeleteBook)
        // protected.PUT("/books/:id", bookController.UpdateBook)

        // группа маршрутов только для авторизованных пользователей
        authenticated := protected.Group("/authenticated")
        authenticated.Use(middleware.RoleMiddleware("authenticated"))
        {
            // Здесь будут маршруты только для авторизованных пользователей
        }

        // ��руппа маршрутов только для администраторов и авторизованных пользователей

		// Группа маршрутов только для администраторов
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			// Здесь будут маршруты только для администраторов
		}
	}

	return r
}
