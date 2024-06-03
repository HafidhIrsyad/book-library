package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	_db "github.com/book-library/app/helper"
	"github.com/book-library/entity/author"
	"gorm.io/gorm"
)

type AuthorRepositoryI interface {
	CreateAuthor(ctx context.Context, trx *gorm.DB, input author.AuthorInput) (err error)
	GetAllAuthors(ctx context.Context, name string) (resp []author.AuthorResponse, err error)
	GetAuthorById(ctx context.Context, id int64, email string) (resp author.AuthorResponse, err error)
	UpdateAuthor(ctx context.Context, trx *gorm.DB, id int64, input author.AuthorInput) (err error)
	DeleteAuthor(ctx context.Context, trx *gorm.DB, id int64) error
}

type AuthorRepository struct {
	conn *gorm.DB
}

func NewAuthorRepository(conn *gorm.DB) AuthorRepositoryI {
	return AuthorRepository{conn: conn}
}

// CreateAuthor implements AuthorRepositoryI.
func (a AuthorRepository) CreateAuthor(ctx context.Context, trx *gorm.DB, input author.AuthorInput) (err error) {
	if trx == nil {
		trx = a.conn.WithContext(ctx)
	}

	now := time.Now()

	input.CreatedAt = now
	input.UpdatedAt = nil

	sql := trx.Table("public" + "." + "tb_author").Create(&input)
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// DeleteAuthor implements AuthorRepositoryI.
func (a AuthorRepository) DeleteAuthor(ctx context.Context, trx *gorm.DB, id int64) error {
	if trx == nil {
		trx = a.conn.WithContext(ctx)
	}

	sql := trx.Table(_db.AuthorTableName).Where("id = ?", id).Delete(&author.AuthorInput{})
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// GetAllAuthors implements AuthorRepositoryI.
func (a AuthorRepository) GetAllAuthors(ctx context.Context, name string) (resp []author.AuthorResponse, err error) {
	query := `SELECT id, name, email, created_at, updated_at FROM ` + _db.AuthorTableName

	params := []interface{}{}
	if name != "" {
		query += ` WHERE lower(name) ilike ?`
		params = append(params, fmt.Sprintf("%%%s%%", strings.ToLower(name)))
	}

	query += ` ORDER BY id ASC`

	sql := a.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// GetAuthorById implements AuthorRepositoryI.
func (a AuthorRepository) GetAuthorById(ctx context.Context, id int64, email string) (resp author.AuthorResponse, err error) {
	params := []interface{}{}
	query := `SELECT id, name, email, created_at, updated_at FROM ` + _db.AuthorTableName

	if id != 0 {
		query += ` WHERE id = ?`
		params = append(params, id)
	}

	if email != "" {
		query += ` WHERE lower(email) = ?`
		params = append(params, email)
	}

	query += ` LIMIT 1`

	sql := a.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// UpdateAuthor implements AuthorRepositoryI.
func (a AuthorRepository) UpdateAuthor(ctx context.Context, trx *gorm.DB, id int64, input author.AuthorInput) (err error) {
	if trx == nil {
		trx = a.conn.WithContext(ctx)
	}

	now := time.Now()
	updateAuthor := map[string]interface{}{
		"name":       input.Name,
		"email":      input.Email,
		"updated_at": &now,
	}

	sql := trx.Table(_db.AuthorTableName).Where("id = ?", id).Updates(updateAuthor)
	if sql.Error != nil {
		return sql.Error
	}

	return err
}
