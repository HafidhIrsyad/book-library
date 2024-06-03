package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	_db "github.com/book-library/app/helper"
	"github.com/book-library/entity/category"
	"gorm.io/gorm"
)

type CategoryRepositoryI interface {
	CreateCategory(ctx context.Context, trx *gorm.DB, input category.CategoryInput) (err error)
	GetAllCategories(ctx context.Context, name string) (resp []category.CategoryResponse, err error)
	GetCategoryById(ctx context.Context, id int64, name string) (resp category.CategoryResponse, err error)
	UpdateCategory(ctx context.Context, trx *gorm.DB, id int64, input category.CategoryInput) (err error)
	DeleteCategory(ctx context.Context, trx *gorm.DB, id int64) error
}

type CategoryRepository struct {
	conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) CategoryRepositoryI {
	return CategoryRepository{conn: conn}
}

// CreateCategory implements CategoryRepositoryI.
func (c CategoryRepository) CreateCategory(ctx context.Context, trx *gorm.DB, input category.CategoryInput) (err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}
	now := time.Now()

	input.CreatedAt = now
	input.UpdatedAt = nil

	sql := trx.Table(_db.CategoryTableName).Create(&input)
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// DeleteCategory implements CategoryRepositoryI.
func (c CategoryRepository) DeleteCategory(ctx context.Context, trx *gorm.DB, id int64) error {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(_db.CategoryTableName).Where("id = ?", id).Delete(&category.CategoryInput{})
	if sql.Error != nil {
		return sql.Error
	}

	return nil
}

// GetAllCategories implements CategoryRepositoryI.
func (c CategoryRepository) GetAllCategories(ctx context.Context, name string) (resp []category.CategoryResponse, err error) {
	query := `SELECT id, name, description, created_at, updated_at FROM tb_category`

	params := []interface{}{}
	if name != "" {
		query += ` WHERE lower(name) ilike ?`
		params = append(params, fmt.Sprintf("%%%s%%", strings.ToLower(name)))
	}

	query += ` ORDER BY id ASC`

	sql := c.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// GetCategoryById implements CategoryRepositoryI.
func (c CategoryRepository) GetCategoryById(ctx context.Context, id int64, name string) (resp category.CategoryResponse, err error) {
	params := []interface{}{}
	query := `SELECT id, name, description, created_at, updated_at FROM ` + _db.CategoryTableName

	if id != 0 {
		query += ` WHERE id = ?`
		params = append(params, id)
	}

	if name != "" {
		query += ` WHERE lower(name) = ?`
		params = append(params, strings.ToLower(name))
	}

	query += ` LIMIT 1`

	sql := c.conn.WithContext(ctx).Raw(query, params...).Scan(&resp)
	if sql.Error != nil {
		return resp, sql.Error
	}

	return resp, err
}

// UpdateCategory implements CategoryRepositoryI.
func (c CategoryRepository) UpdateCategory(ctx context.Context, trx *gorm.DB, id int64, input category.CategoryInput) (err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	now := time.Now()
	updateCategory := map[string]interface{}{
		"name":        input.Name,
		"description": input.Description,
		"updated_at":  &now,
	}

	sql := trx.Table(_db.CategoryTableName).Where("id = ?", id).Updates(updateCategory)
	if sql.Error != nil {
		return sql.Error
	}

	return err
}
