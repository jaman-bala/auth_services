package repositories


import (
	"AuthApplications/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"errors"
)


type BookRepository interface {
	Create(book *models.Book) error
	FindByID(id uuid.UUID) (*models.Book, error)
	FindAll() ([]models.Book, error)
	FindByGenre(genre string) ([]models.Book, error)
	FindByAuthor(author uuid.UUID) ([]models.Book, error)
	Search(query string) ([]models.Book, error)
	Patch(book *models.Book) error
	DeleteByID(id uuid.UUID) error
}

type bookRepository struct {
    db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}

}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}


func (r *bookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) FindByID(id uuid.UUID) (*models.Book, error) {
    var bookID models.Book
    err := r.db.First(&bookID, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("book not found")
        }
        return nil, err
    }
    return &bookID, nil
}


func (r *bookRepository) FindByGenre(genre string) ([]models.Book, error) {
	var bookGenre []models.Book
    err := r.db.Where("genre = ?", genre).Find(&bookGenre).Error
    if err != nil {
        return nil, err
    }
    return bookGenre, nil
}

func (r *bookRepository) FindByAuthor(author uuid.UUID) ([]models.Book, error) {
    var bookAuthor []models.Book
    err := r.db.Where("author = ?", author).Find(&bookAuthor).Error
    if err != nil {
        return nil, err
    }
    return bookAuthor, nil
}

func (r *bookRepository) Search(query string) ([]models.Book, error) {
	var books []models.Book
	// Преобразуем поисковый запрос в нижний регистр для регистронезависимого поиска
	queryLower := strings.ToLower(query)

	// Выполняем поиск по нескольким полям: название, автор, жанр
	err := r.db.Where("lower(title) LIKE ? OR lower(author) LIKE ? OR lower(genre) LIKE ?",
		"%"+queryLower+"%", "%"+queryLower+"%", "%"+queryLower+"%").Find(&books).Error

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) Patch(bookPatch *models.Book) error {
    return r.db.Model(&models.Book{}).Where("id = ?", bookPatch.ID).Updates(bookPatch).Error
}


func (r *bookRepository) DeleteByID(id uuid.UUID) error {
    result := r.db.Delete(&models.Book{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("book not found")
    }
    return nil
}