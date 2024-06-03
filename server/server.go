package server

import (
	"github.com/book-library/app/delivery"
	"github.com/book-library/app/http"
	"github.com/book-library/app/repository"
	"github.com/book-library/app/usecase"
	"github.com/go-chi/chi/v5"
)

func Start() {
	SetConfig(".", ".env")

	dbConn := DBConnection()

	// Repository
	bookRepo := repository.NewBookLibraryRepository(dbConn)
	transactionRepo := repository.NewTransactionRepository(dbConn)
	authorRepo := repository.NewAuthorRepository(dbConn)
	categoryRepo := repository.NewCategoryRepository(dbConn)

	// Usecase
	bookUC := usecase.NewbookLibraryService(bookRepo, transactionRepo, authorRepo, categoryRepo)
	authorUC := usecase.NewAuthorService(authorRepo, transactionRepo, bookRepo)
	categoryUC := usecase.NewCategoryService(categoryRepo, transactionRepo, bookRepo)

	// Handler
	bookHandler := delivery.NewBookHandler(bookUC)
	authorHandler := delivery.NewAuthorHandler(authorUC)
	categoryHandler := delivery.NewCategoryHandler(categoryUC)

	r := chi.NewRouter()
	Set(r)

	// Router
	http.BookPath(r, bookHandler)
	http.AuthorPath(r, authorHandler)
	http.CategoryPath(r, categoryHandler)

	startServerWithGracefulShutdown(r)
}
