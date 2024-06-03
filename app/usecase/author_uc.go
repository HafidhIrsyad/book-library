package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	_track "github.com/book-library/app/helper"
	_r "github.com/book-library/app/repository"
	"github.com/book-library/entity/author"
	_l "github.com/rs/zerolog/log"
)

type AuthorServiceI interface {
	CreateAuthor(ctx context.Context, input author.AuthorInput) (err error)
	UpdateAuthor(ctx context.Context, id int64, input author.AuthorInput) (err error)
	GetAuthorByID(ctx context.Context, id int64) (resp author.AuthorResponse, err error)
	GetAllAuthors(ctx context.Context, name string) (resp []author.AuthorResponse, err error)
	DeleteAuthorByID(ctx context.Context, id int64) (err error)
}

type AuthorService struct {
	authorRepo _r.AuthorRepositoryI
	trRepo     _r.TransactionRepositoryI
	bookRepo   _r.BookLibraryRepositoryI
}

func NewAuthorService(authorRepo _r.AuthorRepositoryI, trRepo _r.TransactionRepositoryI, bookRepo _r.BookLibraryRepositoryI) AuthorServiceI {
	return AuthorService{
		authorRepo: authorRepo,
		trRepo:     trRepo,
		bookRepo:   bookRepo,
	}
}

// CreateAuthor implements AuthorServiceI.
func (a AuthorService) CreateAuthor(ctx context.Context, input author.AuthorInput) (err error) {
	defer _track.TimeTrack(time.Now(), "CreateAuthor")

	if a.validationInput(input); err != nil {
		_l.Error().Err(err).Msg("a.validationInput got an error on AuthorService.CreateAuthor")
		return err
	}

	trx := a.trRepo.BeginTransaction(ctx)

	err = a.authorRepo.CreateAuthor(ctx, trx, input)
	if err != nil {
		_l.Error().Err(err).Msg("a.authorRepo.CreateAuthor got an error on AuthorService.CreateAuthor")
		a.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	a.trRepo.CommitTransaction(ctx, trx)

	return err
}

// DeleteAuthorByID implements AuthorServiceI.
func (a AuthorService) DeleteAuthorByID(ctx context.Context, id int64) (err error) {
	defer _track.TimeTrack(time.Now(), "DeleteAuthorByID")
	_log := _l.Ctx(ctx)

	bookByID, err := a.bookRepo.GetBookLibraryById(ctx, 0, id, 0)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.GetBookLibraryById got an error on AuthorService.DeleteAuthorByID")
		return err
	}

	if bookByID.ID != 0 {
		_log.Error().Msgf("There is book(%s) using this author and delete book(%s) first before delete author", bookByID.Title, bookByID.Title)
		return fmt.Errorf("There is book(%s) using this author and delete the book(%s) first before delete author", bookByID.Title, bookByID.Title)
	}

	trx := a.trRepo.BeginTransaction(ctx)

	err = a.authorRepo.DeleteAuthor(ctx, trx, id)
	if err != nil {
		_log.Error().Err(err).Msg("c.authorRepo.DeleteAuthor got an error on AuthorService.DeleteAuthorByID")
		a.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	a.trRepo.CommitTransaction(ctx, trx)

	return err
}

// GetAllAuthors implements AuthorServiceI.
func (a AuthorService) GetAllAuthors(ctx context.Context, name string) (resp []author.AuthorResponse, err error) {
	defer _track.TimeTrack(time.Now(), "GetAllAuthors")
	_log := _l.Ctx(ctx)

	categories, err := a.authorRepo.GetAllAuthors(ctx, name)
	if err != nil {
		_log.Error().Err(err).Msg("c.authorRepo.GetAllAuthors got an error on AuthorService.GetAllAuthors")
		return resp, err
	}

	return categories, err
}

// GetAuthorByID implements AuthorServiceI.
func (a AuthorService) GetAuthorByID(ctx context.Context, id int64) (resp author.AuthorResponse, err error) {
	defer _track.TimeTrack(time.Now(), "GetAuthorByID")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("AuthorID cannot be nol on AuthorService.GetAuthorByID")
		return resp, errors.New("AuthorID cannot be nol")
	}

	authorById, err := a.authorRepo.GetAuthorById(ctx, id, "")
	if err != nil {
		_log.Error().Err(err).Msg("c.authorRepo.GetAuthorById got an error on AuthorService.GetAuthorByID")
		return resp, err
	}

	if authorById.ID == 0 {
		_log.Error().Err(err).Msg("Author not found on AuthorService.GetAuthorByID")
		return resp, errors.New("Author not found")
	}

	return authorById, err
}

// UpdateAuthor implements AuthorServiceI.
func (a AuthorService) UpdateAuthor(ctx context.Context, id int64, input author.AuthorInput) (err error) {
	defer _track.TimeTrack(time.Now(), "UpdateAuthor")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("AuthorID cannot be nol on AuthorService.UpdateAuthor")
		return errors.New("AuthorID cannot be nol")
	}

	authorById, err := a.authorRepo.GetAuthorById(ctx, id, "")
	if err != nil {
		_log.Error().Err(err).Msg("c.GetAuthorByID.GetAuthorById got an error on AuthorService.UpdateAuthor")
		return err
	}

	if authorById.ID == 0 {
		_log.Error().Msg("Author not found on AuthorService.UpdateAuthor")
		return errors.New("Author not found")
	}

	if input.Name == "" {
		input.Name = authorById.Name
	}

	if input.Email == "" {
		input.Email = authorById.Email
	}

	trx := a.trRepo.BeginTransaction(ctx)

	err = a.authorRepo.UpdateAuthor(ctx, trx, id, input)
	if err != nil {
		_log.Error().Err(err).Msg("c.authorRepo.UpdateAuthor got an error on AuthorService.UpdateAuthor")
		a.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	a.trRepo.CommitTransaction(ctx, trx)

	return err
}

func (a AuthorService) validationInput(input author.AuthorInput) (err error) {
	if input.Name == "" {
		return errors.New("Name can not be empty")
	}

	if input.Email == "" {
		return errors.New("Email can not be empty")
	}

	return nil
}
