package repository

import (
	"context"
	"time"

	_db "github.com/book-library/app/helper"
	"github.com/book-library/entity/book"
	"gorm.io/gorm"
)

type BookLibraryRepositoryI interface {
	CreateBookLibrary(ctx context.Context, trx *gorm.DB, input book.BookInput) (err error)
	GetAllBookLibraries(ctx context.Context, search string) (resp []book.BookResponse, err error)
	GetBookLibraryById(ctx context.Context, id, authorID, categoryID int64) (resp book.BookResponse, err error)
	UpdateBookLibrary(ctx context.Context, trx *gorm.DB, id int64, input book.BookInput) (rerr error)
	DeleteBookLibrary(ctx context.Context, trx *gorm.DB, id int64) error
}

type BookLibraryRepository struct {
	conn *gorm.DB
}

func NewBookLibraryRepository(conn *gorm.DB) BookLibraryRepositoryI {
	return BookLibraryRepository{conn: conn}
}

// CreateBookLibrary implements BookLibraryRepositoryI.
func (b BookLibraryRepository) CreateBookLibrary(ctx context.Context, trx *gorm.DB, input book.BookInput) (err error) {
	if trx == nil {
		trx = b.conn.WithContext(ctx)
	}

	now := time.Now()

	input.CreatedAt = now
	input.UpdatedAt = nil
	sql := trx.Table(_db.BookTableName).Create(&input)
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// DeleteBookLibrary implements BookLibraryRepositoryI.
func (b BookLibraryRepository) DeleteBookLibrary(ctx context.Context, trx *gorm.DB, id int64) error {
	if trx == nil {
		trx = b.conn.WithContext(ctx)
	}

	sql := trx.Table(_db.BookTableName).Where("id = ?", id).Delete(&book.BookInput{})
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// GetAllBookLibrary implements BookLibraryRepositoryI.
func (b BookLibraryRepository) GetAllBookLibraries(ctx context.Context, search string) (resp []book.BookResponse, err error) {
	query := `
		SELECT
			tbb.id, tbb.title, tbb.description, tbb.isbn, tbb.published_flag, 
			tba.id as author_id, tba.name as author_name, tba.email as author_email,
			tbc.id as category_id, tbc.name as category_name, tbc.description as category_description
		FROM 
			tb_book tbb
		LEFT JOIN 
			tb_category tbc on tbb.category_id = tbc.id 
		LEFT JOIN 
			tb_author tba on tbb.author_id = tba.id 
		WHERE
			tbb.published_flag = true
	`
	params := []interface{}{}

	if search != "" {
		query += ` AND tbb.isbn = ? OR tbb.title = ? OR tba.name = ?`
		params = append(params, search, search, search)
	}

	sql := b.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// GetBookLibraryById implements BookLibraryRepositoryI.
func (b BookLibraryRepository) GetBookLibraryById(ctx context.Context, id, authorID, categoryID int64) (resp book.BookResponse, err error) {
	query := `
		SELECT
			tbb.id, tbb.title, tbb.isbn, tbb.description,tbb.published_flag, tbb.author_id, tbb.category_id
		FROM 
			tb_book tbb
	`

	params := []interface{}{}
	if id != 0 {
		query += ` WHERE tbb.id = ?`
		params = append(params, id)
	}

	if authorID != 0 {
		query += ` WHERE tbb.author_id = ?`
		params = append(params, authorID)
	}

	if categoryID != 0 {
		query += ` WHERE tbb.category_id = ?`
		params = append(params, categoryID)
	}

	sql := b.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// UpdateBookLibrary implements BookLibraryRepositoryI.
func (b BookLibraryRepository) UpdateBookLibrary(ctx context.Context, trx *gorm.DB, id int64, input book.BookInput) (err error) {
	if trx == nil {
		trx = b.conn.WithContext(ctx)
	}

	now := time.Now()
	updateBookLibrary := map[string]interface{}{
		"title":          input.Title,
		"isbn":           input.ISBN,
		"description":    input.Description,
		"published_flag": input.PublishedFlag,
		"author_id":      input.AuthorID,
		"category_id":    input.CategoryID,
		"updated_at":     &now,
	}

	sql := trx.Table(_db.BookTableName).Where("id = ?", id).Updates(updateBookLibrary)
	if sql.Error != nil {
		return sql.Error
	}

	return err
}
