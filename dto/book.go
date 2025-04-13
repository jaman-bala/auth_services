package dto

import (
	"github.com/google/uuid"
)

type BookRequest struct {
	Title       string    `json:"title"`
	AuthorID    uuid.UUID `json:"author_id"`
	Description string    `json:"description"`
	ISBN        string    `json:"isbn"`
	PublishYear int       `json:"publish_year"`
	CoverURL    string    `json:"cover_url"`
	FileURL     string    `json:"file_url"`
	Genre       string    `json:"genre"`
	Language    string    `json:"language"`
	PageCount   int       `json:"page_count"`
}


type BookResponse struct {
	ID          uuid.UUID `json:"id"`
	Title string `json:"title"`
	AuthorID uuid.UUID `json:"author_id"`
	Description string `json:"description"`
	ISBN string `json:"isbn"`
	PublishYear int `json:"publish_year"`
	CoverURL string `json:"cover_url"`
	FileURL string `json:"file_url"`
	Genre string `json:"genre"`
	Language string `json:"language"`
	PageCount int `json:"page_count"`
}


type PatchBookRequest struct {
	Title       *string    `json:"title,omitempty"`
	AuthorID    *uuid.UUID `json:"author_id,omitempty"`
	Description *string    `json:"description,omitempty"`
	ISBN        *string    `json:"isbn,omitempty"`
	PublishYear *int       `json:"publish_year,omitempty"`
	CoverURL    *string    `json:"cover_url,omitempty"`
	FileURL     *string    `json:"file_url,omitempty"`
	Genre       *string    `json:"genre,omitempty"`
	Language    *string    `json:"language,omitempty"`
	PageCount   *int       `json:"page_count,omitempty"`
}