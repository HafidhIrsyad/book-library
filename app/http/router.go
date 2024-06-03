package http

import (
	"github.com/book-library/app/delivery"
	"github.com/go-chi/chi/v5"
)

func BookPath(r *chi.Mux, bh delivery.BookHandler) {
	r.Route("/api/v1/book/", func(r chi.Router) {
		r.Post("/create", bh.CreateBook)
		r.Put("/update/{id}", bh.UpdateBook)
		r.Get("/all", bh.GetBooks)
		r.Get("/{id}", bh.GetBookById)
		r.Delete("/{id}", bh.DeleteBookyByID)
	})
}

func AuthorPath(r *chi.Mux, ah delivery.AuthorHandler) {
	r.Route("/api/v1/author", func(r chi.Router) {
		r.Post("/create", ah.CreateAuthor)
		r.Put("/update/{id}", ah.UpdateAuthor)
		r.Get("/all", ah.GetAuthors)
		r.Get("/{id}", ah.GetAuhtorById)
		r.Delete("/{id}", ah.DeleteAuthorByID)
	})
}

func CategoryPath(r *chi.Mux, ch delivery.CategoryHandler) {
	r.Route("/api/v1/category", func(r chi.Router) {
		r.Post("/create", ch.CreateCategory)
		r.Put("/update/{id}", ch.UpdateCategory)
		r.Get("/all", ch.GetCategories)
		r.Get("/{id}", ch.GetCategoryById)
		r.Delete("/{id}", ch.DeleteCategoryByID)
	})
}
