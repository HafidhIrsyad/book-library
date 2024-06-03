package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	_track "github.com/book-library/app/helper"
	_r "github.com/book-library/app/repository"
	"github.com/book-library/entity/category"
	_l "github.com/rs/zerolog/log"
)

type CategoryServiceI interface {
	CreateCategory(ctx context.Context, input category.CategoryInput) (err error)
	UpdateCategory(ctx context.Context, id int64, input category.CategoryInput) (err error)
	GetCategoryByID(ctx context.Context, id int64) (resp category.CategoryResponse, err error)
	GetAllCategories(ctx context.Context, name string) (resp []category.CategoryResponse, err error)
	DeleteCategoryByID(ctx context.Context, id int64) (err error)
}

type CategoryService struct {
	categoryRepo _r.CategoryRepositoryI
	trRepo       _r.TransactionRepositoryI
	bookRepo     _r.BookLibraryRepositoryI
}

func NewCategoryService(categoryRepo _r.CategoryRepositoryI, trRepo _r.TransactionRepositoryI, bookRepo _r.BookLibraryRepositoryI) CategoryServiceI {
	return CategoryService{
		categoryRepo: categoryRepo,
		trRepo:       trRepo,
		bookRepo:     bookRepo,
	}
}

// CreateCategory implements CategoryServiceI.
func (c CategoryService) CreateCategory(ctx context.Context, input category.CategoryInput) (err error) {
	defer _track.TimeTrack(time.Now(), "CreateCategoryUC")
	_log := _l.Ctx(ctx)

	if c.validationInput(input); err != nil {
		_log.Error().Err(err).Msg("c.validationInput got an error on CategoryService.CreateCategory")
		return err
	}

	byID, err := c.categoryRepo.GetCategoryById(ctx, 0, strings.ToLower(input.Name))
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.GetCategoryById got an error on CategoryService.CreateCategory")
		return err
	}

	if byID.ID != 0 {
		_log.Error().Err(err).Msgf("Category %s is already exist", byID.Name)
		return fmt.Errorf("Category %s is already exist", byID.Name)
	}

	trx := c.trRepo.BeginTransaction(ctx)

	err = c.categoryRepo.CreateCategory(ctx, trx, input)
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.CreateCategory got an error on CategoryService.CreateCategory")
		c.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	c.trRepo.CommitTransaction(ctx, trx)

	return err
}

// DeleteCategoryByID implements CategoryServiceI.
func (c CategoryService) DeleteCategoryByID(ctx context.Context, id int64) (err error) {
	defer _track.TimeTrack(time.Now(), "DeleteCategoryByIDUC")
	_log := _l.Ctx(ctx)

	bookByID, err := c.bookRepo.GetBookLibraryById(ctx, 0, 0, id)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.GetBookLibraryById got an error on CategoryService.DeleteCategoryByID")
		return err
	}

	if bookByID.ID != 0 {
		_log.Error().Msgf("There is book(%s) using this category and delete book(%s) first before delete category", bookByID.Title, bookByID.Title)
		return fmt.Errorf("There is book(%s) using this category and delete the book(%s) first before delete category", bookByID.Title, bookByID.Title)
	}

	trx := c.trRepo.BeginTransaction(ctx)

	err = c.categoryRepo.DeleteCategory(ctx, trx, id)
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.DeleteCategory got an error on CategoryService.DeleteCategoryByID")
		c.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	c.trRepo.CommitTransaction(ctx, trx)

	return err
}

// GetAllCategories implements CategoryServiceI.
func (c CategoryService) GetAllCategories(ctx context.Context, name string) (resp []category.CategoryResponse, err error) {
	defer _track.TimeTrack(time.Now(), "GetAllCategoriesUC")
	_log := _l.Ctx(ctx)

	categories, err := c.categoryRepo.GetAllCategories(ctx, name)
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.GetAllCategories got an error on CategoryService.GetAllCategories")
		return resp, err
	}

	return categories, err
}

// GetCategoryByID implements CategoryServiceI.
func (c CategoryService) GetCategoryByID(ctx context.Context, id int64) (resp category.CategoryResponse, err error) {
	defer _track.TimeTrack(time.Now(), "GetCategoryByIDUC")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("CategoryID cannot be nol on CategoryService.GetCategoryByID")
		return resp, errors.New("CategoryID cannot be nol")
	}

	catById, err := c.categoryRepo.GetCategoryById(ctx, id, "")
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.GetCategoryById got an error on CategoryService.GetCategoryByID")
		return resp, err
	}

	if catById.ID == 0 {
		_log.Error().Err(err).Msg("Category not found on CategoryService.GetCategoryByID")
		return resp, errors.New("Category not found")
	}

	return catById, err
}

// UpdateCategory implements CategoryServiceI.
func (c CategoryService) UpdateCategory(ctx context.Context, id int64, input category.CategoryInput) (err error) {
	defer _track.TimeTrack(time.Now(), "UpdateCategoryUC")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("CategoryID cannot be nol on CategoryService.UpdateCategory")
		return errors.New("CategoryID cannot be nol")
	}

	catById, err := c.categoryRepo.GetCategoryById(ctx, id, "")
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.GetCategoryById got an error on CategoryService.UpdateCategory")
		return err
	}

	if catById.ID == 0 {
		_log.Error().Msg("Category not found on CategoryService.UpdateCategory")
		return errors.New("Category not found")
	}

	if input.Description == "" {
		input.Description = catById.Description
	}

	if input.Name == "" {
		input.Description = catById.Name
	}

	trx := c.trRepo.BeginTransaction(ctx)

	err = c.categoryRepo.UpdateCategory(ctx, trx, id, input)
	if err != nil {
		_log.Error().Err(err).Msg("c.categoryRepo.UpdateCategory got an error on CategoryService.UpdateCategory")
		c.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	c.trRepo.CommitTransaction(ctx, trx)

	return err
}

func (c CategoryService) validationInput(input category.CategoryInput) (err error) {
	if input.Name == "" {
		return errors.New("Name can not be empty")
	}

	if input.Description == "" {
		return errors.New("Description can not be empty")
	}

	return nil
}
