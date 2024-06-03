package book

import (
	"time"

	"github.com/book-library/entity/author"
	"github.com/book-library/entity/category"
)

type (
	BookInput struct {
		Title         string     `json:"title"`
		AuthorID      int64      `json:"author_id"`
		Description   string     `json:"description"`
		ISBN          string     `json:"isbn"`
		PublishedFlag *bool      `json:"published_flag"`
		CategoryID    int64      `json:"category_id"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
	}

	BookResponse struct {
		ID                  int64      `json:"id"`
		Title               string     `json:"title"`
		BoookDescription    string     `json:"description"`
		AuthorID            int64      `json:"author_id"`
		AuthorName          string     `json:"author_name"`
		AuthorEmail         string     `json:"author_email"`
		CategoryID          int64      `json:"category_id"`
		CategoryName        string     `json:"category_name"`
		CategoryDescription string     `json:"category_description"`
		ISBN                string     `json:"isbn"`
		PublishedFlag       bool       `json:"published_flag"`
		CreatedAt           time.Time  `json:"created_at"`
		UpdatedAt           *time.Time `json:"updated_at"`
	}

	BookResponseDetail struct {
		ID            int64                         `json:"id"`
		Title         string                        `json:"title"`
		Description   string                        `json:"description"`
		Author        author.AuthorResponseJoin     `json:"author"`
		Category      category.CategoryResponseJoin `json:"category"`
		ISBN          string                        `json:"isbn"`
		PublishedFlag bool                          `json:"published_flag"`
		CreatedAt     time.Time                     `json:"created_at"`
		UpdatedAt     *time.Time                    `json:"updated_at"`
	}
)
