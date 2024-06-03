package usecase

import (
	"context"
	"errors"
	"time"

	_track "github.com/book-library/app/helper"
	_r "github.com/book-library/app/repository"
	"github.com/book-library/entity/author"
	"github.com/book-library/entity/book"
	"github.com/book-library/entity/category"
	_l "github.com/rs/zerolog/log"
)

type BookLibraryServiceI interface {
	CreateBook(ctx context.Context, input book.BookInput) (err error)
	UpdateBook(ctx context.Context, id int64, input book.BookInput) (err error)
	GetBookByID(ctx context.Context, id int64) (resp book.BookResponseDetail, err error)
	GetAllBooks(ctx context.Context, search string) (resp []book.BookResponseDetail, err error)
	DeleteBookByID(ctx context.Context, id int64) (err error)
}

type BookLibraryService struct {
	bookRepo     _r.BookLibraryRepositoryI
	trRepo       _r.TransactionRepositoryI
	authorRepo   _r.AuthorRepositoryI
	categoryRepo _r.CategoryRepositoryI
}

func NewbookLibraryService(bookRepo _r.BookLibraryRepositoryI, trRepo _r.TransactionRepositoryI, authorRepo _r.AuthorRepositoryI, categoryRepo _r.CategoryRepositoryI) BookLibraryServiceI {
	return BookLibraryService{
		bookRepo:     bookRepo,
		trRepo:       trRepo,
		authorRepo:   authorRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateBook implements BookLibraryServiceI.
func (b BookLibraryService) CreateBook(ctx context.Context, input book.BookInput) (err error) {
	now := time.Now()
	defer _track.TimeTrack(now, "CreateBookUC")

	if b.validationInput(input); err != nil {
		_l.Error().Err(err).Msg("b.validationInput got an error on BookLibraryService.CreateBook")
		return err
	}

	trx := b.trRepo.BeginTransaction(ctx)

	err = b.bookRepo.CreateBookLibrary(ctx, trx, input)
	if err != nil {
		_l.Error().Err(err).Msg("b.repo.CreateBookLibrary got an error on BookLibraryService.CreateBook")
		b.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	b.trRepo.CommitTransaction(ctx, trx)

	return err
}

// UpdateBook implements BookLibraryServiceI.
func (b BookLibraryService) UpdateBook(ctx context.Context, id int64, input book.BookInput) (err error) {
	defer _track.TimeTrack(time.Now(), "UpdateBook")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("BookID cannot be nol on BookLibraryService.UpdateBook")
		return errors.New("BookID cannot be nol")
	}

	bookById, err := b.bookRepo.GetBookLibraryById(ctx, id, 0, 0)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.GetBookLibraryById got an error on BookLibraryService.UpdateBook")
		return err
	}

	if bookById.ID == 0 {
		_log.Error().Msg("Book not found on BookLibraryService.UpdateBook")
		return errors.New("Book not found")
	}

	if input.Title == "" {
		input.Title = bookById.Title
	}

	if input.AuthorID == 0 {
		input.AuthorID = bookById.AuthorID
	}

	if input.CategoryID == 0 {
		input.CategoryID = bookById.CategoryID
	}

	if input.ISBN == "" {
		input.ISBN = bookById.ISBN
	}

	if input.Description == "" {
		input.Description = bookById.BoookDescription
	}

	if input.PublishedFlag == nil {
		input.PublishedFlag = &bookById.PublishedFlag
	}

	trx := b.trRepo.BeginTransaction(ctx)

	err = b.bookRepo.UpdateBookLibrary(ctx, trx, id, input)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.UpdateBookLibrary got an error on BookLibraryService.UpdateBook")
		b.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	b.trRepo.CommitTransaction(ctx, trx)

	return err
}

// GetAllBooks implements BookLibraryServiceI.
func (b BookLibraryService) GetAllBooks(ctx context.Context, search string) (resp []book.BookResponseDetail, err error) {
	defer _track.TimeTrack(time.Now(), "GetAllBooks")
	_log := _l.Ctx(ctx)

	books, err := b.bookRepo.GetAllBookLibraries(ctx, search)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.GetAllBookLibraries got an error on BookLibraryService.GetAllBooks")
		return resp, err
	}

	booksResp := []book.BookResponseDetail{}
	for _, v := range books {
		bookResp := book.BookResponseDetail{
			ID:            v.ID,
			Title:         v.Title,
			Description:   v.BoookDescription,
			ISBN:          v.ISBN,
			PublishedFlag: v.PublishedFlag,
			Author: author.AuthorResponseJoin{
				ID:    v.AuthorID,
				Name:  v.AuthorName,
				Email: v.AuthorEmail,
			},
			Category: category.CategoryResponseJoin{
				ID:          v.CategoryID,
				Name:        v.CategoryName,
				Description: v.CategoryDescription,
			},
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		booksResp = append(booksResp, bookResp)
	}

	return booksResp, err
}

// GetBookByID implements BookLibraryServiceI.
func (b BookLibraryService) GetBookByID(ctx context.Context, id int64) (resp book.BookResponseDetail, err error) {
	defer _track.TimeTrack(time.Now(), "GetAuthorByID")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("BookID cannot be nol on BookLibraryService.GetBookByID")
		return resp, errors.New("BookID cannot be nol")
	}

	bookById, err := b.bookRepo.GetBookLibraryById(ctx, id, 0, 0)
	if err != nil {
		_log.Error().Err(err).Msg("b.bookRepo.GetBookLibraryById got an error on BookLibraryService.GetBookByID")
		return resp, err
	}

	if bookById.ID == 0 {
		_log.Error().Err(err).Msg("Book not found on BookLibraryService.GetBookByID")
		return resp, errors.New("Book not found")
	}

	authorById, err := b.authorRepo.GetAuthorById(ctx, bookById.AuthorID, "")
	if err != nil {
		_log.Error().Err(err).Msg("b.authorRepo.GetAuthorById got an error on BookLibraryService.GetBookByID")
		return resp, err
	}

	if authorById.ID == 0 {
		_log.Error().Err(err).Msg("Author not found on BookLibraryService.GetBookByID")
		return resp, errors.New("Author not found")
	}

	categoryById, err := b.categoryRepo.GetCategoryById(ctx, bookById.CategoryID, "")
	if err != nil {
		_log.Error().Err(err).Msg("b.categoryRepo.GetCategoryById got an error on BookLibraryService.GetBookByID")
		return resp, err
	}

	if categoryById.ID == 0 {
		_log.Error().Err(err).Msg("Category not found on BookLibraryService.GetBookByID")
		return resp, errors.New("Category not found")
	}

	resp = book.BookResponseDetail{
		ID:            bookById.ID,
		Title:         bookById.Title,
		Description:   bookById.BoookDescription,
		ISBN:          bookById.ISBN,
		PublishedFlag: bookById.PublishedFlag,
		Author: author.AuthorResponseJoin{
			ID:    authorById.ID,
			Name:  authorById.Name,
			Email: authorById.Email,
		},
		Category: category.CategoryResponseJoin{
			ID:          categoryById.ID,
			Name:        categoryById.Name,
			Description: categoryById.Description,
		},
		CreatedAt: bookById.CreatedAt,
		UpdatedAt: bookById.UpdatedAt,
	}

	return resp, err
}

// DeleteBookByID implements BookLibraryServiceI.
func (b BookLibraryService) DeleteBookByID(ctx context.Context, id int64) (err error) {
	defer _track.TimeTrack(time.Now(), "DeleteBookByID")
	_log := _l.Ctx(ctx)

	if id == 0 {
		_log.Error().Msg("BookID cannot be nol on BookLibraryService.DeleteBookByID")
		return errors.New("BookID cannot be nol")
	}

	trx := b.trRepo.BeginTransaction(ctx)

	err = b.bookRepo.DeleteBookLibrary(ctx, trx, id)
	if err != nil {
		_log.Error().Err(err).Msg("c.bookRepo.DeleteBookLibrary got an error on BookLibraryService.DeleteBookByID")
		b.trRepo.RollBackTransaction(ctx, trx)
		return err
	}

	b.trRepo.CommitTransaction(ctx, trx)

	return err
}

func (b BookLibraryService) validationInput(input book.BookInput) (err error) {
	if input.AuthorID == 0 {
		return errors.New("AuthorID can not be zero / 0")
	}

	if input.CategoryID == 0 {
		return errors.New("CategoryID can not be zero / 0")
	}

	if input.Title == "" {
		return errors.New("Title can not be empty")
	}

	if input.Description == "" {
		return errors.New("Description can not be empty")
	}

	if input.ISBN == "" {
		return errors.New("ISBN can not be empty")
	}

	return nil
}
