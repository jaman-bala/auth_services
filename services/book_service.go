package services

import (
	"AuthApplications/dto"
	"AuthApplications/models"
	"AuthApplications/repositories"
	"errors"

	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(req dto.BookRequest) (*models.Book, error)
	GetAllBook() ([]*dto.BookResponse, error)
	GetByID(id uuid.UUID) (*dto.BookResponse, error)
	FindByGenre(genre string) ([]*dto.BookResponse, error)
    Search(query string) ([]*dto.BookResponse, error)
    PatchBook(bookID uuid.UUID, req dto.PatchBookRequest) (*dto.BookResponse, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) CreateBook(req dto.BookRequest) (*models.Book, error) {
	newBook := &models.Book{
		Title:       req.Title,
		AuthorID:    req.AuthorID,
		Description: req.Description,
		ISBN:        req.ISBN,
		PublishYear: req.PublishYear,
		CoverURL:    req.CoverURL,
		FileURL:     req.FileURL,
		Genre:       req.Genre,
		Language:    req.Language,
		PageCount:   req.PageCount,
	}

	err := s.bookRepo.Create(newBook)
	if err != nil {
		return nil, err
	}

	return newBook, nil
}

func (s *bookService) GetAllBook() ([]*dto.BookResponse, error) {
	books, err := s.bookRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var bookResponses []*dto.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, &dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			AuthorID:    book.AuthorID,
			Description: book.Description,
			ISBN:        book.ISBN,
			PublishYear: book.PublishYear,
			CoverURL:    book.CoverURL,
			FileURL:     book.FileURL,
			Genre:       book.Genre,
			Language:    book.Language,
			PageCount:   book.PageCount,
		})
	}

	return bookResponses, nil
}

func (s *bookService) GetByID(id uuid.UUID) (*dto.BookResponse, error) {
	book, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	return &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		Description: book.Description,
		ISBN:        book.ISBN,
		PublishYear: book.PublishYear,
		CoverURL:    book.CoverURL,
		FileURL:     book.FileURL,
		Genre:       book.Genre,
		Language:    book.Language,
		PageCount:   book.PageCount,
	}, nil
}

func (s *bookService) FindByGenre(genre string) ([]*dto.BookResponse, error) {
	books, err := s.bookRepo.FindByGenre(genre)
	if err != nil {
		return nil, err
	}
    if books == nil {
        return nil, errors.New("no books found with this genre")
    }

	var bookResponses []*dto.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, &dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			AuthorID:    book.AuthorID,
			Description: book.Description,
			ISBN:        book.ISBN,
			PublishYear: book.PublishYear,
			CoverURL:    book.CoverURL,
			FileURL:     book.FileURL,
			Genre:       book.Genre,
			Language:    book.Language,
			PageCount:   book.PageCount,
		})
	}

	return bookResponses, nil
}


func (s *bookService) Search(query string) ([]*dto.BookResponse, error) {
    books, err := s.bookRepo.Search(query)
    if err != nil {
        return nil, err
    }
    if len(books) == 0 {
        return []*dto.BookResponse{}, errors.New("no books found matching this query")
    }
    var bookResponses []*dto.BookResponse
    for _, book := range books {
        bookResponses = append(bookResponses, &dto.BookResponse{
            ID:          book.ID,
            Title:       book.Title,
            AuthorID:    book.AuthorID,
            Description: book.Description,
            ISBN:        book.ISBN,
            PublishYear: book.PublishYear,
            CoverURL:    book.CoverURL,
            FileURL:     book.FileURL,
            Genre:       book.Genre,
            Language:    book.Language,
            PageCount:   book.PageCount,
        })
    }
    return bookResponses, nil
}

func (s *bookService) PatchBook(bookID uuid.UUID, req dto.PatchBookRequest) (*dto.BookResponse, error) {
	book, err := s.bookRepo.FindByID(bookID)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	if req.Title != nil {
		book.Title = *req.Title
	}
	if req.AuthorID != nil {
		book.AuthorID = *req.AuthorID
	}
	if req.Description != nil {
		book.Description = *req.Description
	}
	if req.ISBN != nil {
		book.ISBN = *req.ISBN
	}
	if req.PublishYear != nil {
		book.PublishYear = *req.PublishYear
	}
	if req.CoverURL != nil {
		book.CoverURL = *req.CoverURL
	}
	if req.FileURL != nil {
		book.FileURL = *req.FileURL
	}
	if req.Genre != nil {
		book.Genre = *req.Genre
	}
	if req.Language != nil {
		book.Language = *req.Language
	}
	if req.PageCount != nil {
		book.PageCount = *req.PageCount
	}
	err = s.bookRepo.Patch(book)
	if err != nil {
		return nil, err
	}
	return &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		AuthorID:    book.AuthorID,
		Description: book.Description,
		ISBN:        book.ISBN,
		PublishYear: book.PublishYear,
		CoverURL:    book.CoverURL,
		FileURL:     book.FileURL,
		Genre:       book.Genre,
		Language:    book.Language,
		PageCount:   book.PageCount,
	}, nil
}