package controllers

import (
	"net/http"

	"AuthApplications/dto"
	"AuthApplications/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// BookController интерфейс для методов книги
type BookController interface {
	CreateBook(c *gin.Context)
    GetAllBooks(c *gin.Context)
    GetByID(c *gin.Context)
    FindByGenre(c *gin.Context)
    
}

// Реализация BookController
type bookController struct {
    bookService services.BookService
}

// Конструктор для создания экземпляра BookController
func NewBookController(bookService services.BookService) BookController {
    return &bookController{
        bookService: bookService,
    }
}

// CreateBook godoc
// @Summary Создание книги
// @Description Создаёт новую книгу
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body dto.BookRequest true "Данные книг"
// @Success 200 {array} dto.UserResponse "Список пользователей"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/books [post]
func (bc *bookController) CreateBook(c *gin.Context) {
    var request dto.BookRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    createdBook, err := bc.bookService.CreateBook(request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Книга успешно создана",
        "book_id": createdBook.ID,
    })
}


// GetAllBooks godoc
// @Summary Получение всех книг
// @Description Возвращает список всех книг
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.BookResponse "Список книг"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/books [get]
func (bc *bookController) GetAllBooks(c *gin.Context) {
    books, err := bc.bookService.GetAllBook()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, books)
}



// GetByID godoc
// @Summary Получение книги по ID
// @Description Возвращает информацию о книги по указанному ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "ID книги"
// @Security BearerAuth
// @Success 200 {object} dto.BookResponse "Книга найден"
// @Failure 400 {object} map[string]string "Некорректный запрос"
// @Failure 404 {object} map[string]string "Книга не найден"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/books/{id} [get]
func (bc *bookController) GetByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID"})
        return
    }

    book, err := bc.bookService.GetByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, book)
}


// FindByGenre godoc
// @Summary Поиск книг по жанру
// @Description Возвращает список книг по указанному жанру
// @Tags Book
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param genre path string true "Жанр книги"
// @Success 200 {array} dto.BookResponse "Список книг"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/books/genre/{genre} [get]
func (bc *bookController) FindByGenre(c *gin.Context) {
    genre := c.Param("genre")

    books, err := bc.bookService.FindByGenre(genre)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, books)
}


