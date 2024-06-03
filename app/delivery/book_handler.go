package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	api "github.com/book-library/app/helper"
	u "github.com/book-library/app/usecase"
	"github.com/book-library/entity/book"
	"github.com/book-library/logger"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type BookHandler struct {
	bookUC u.BookLibraryServiceI
}

func NewBookHandler(bookUC u.BookLibraryServiceI) BookHandler {
	return BookHandler{
		bookUC: bookUC,
	}
}

func (h BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())

	var input book.BookInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on BookHandler.CreateBook"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.bookUC.CreateBook(ctx, input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.bookUC.CreateBook got an error on BookHandler.CreateBook"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to Created Book", Code: http.StatusOK, Success: true})
}

func (h BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	var input book.BookInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on BookHandler.UpdateBook"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.bookUC.UpdateBook(ctx, int64(idInt), input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.bookUC.UpdateBook got an error on BookHandler.UpdateBook"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to UpdateBook", Code: http.StatusOK, Success: true})
}

func (h BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	catById, err := h.bookUC.GetBookByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.bookUC.GetBookByID got an error on BookHandler.GetBookById"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetBookById", Code: http.StatusOK, Success: true}, Data: catById})
}

func (h BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	name := r.URL.Query().Get("name")

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: name})

	books, err := h.bookUC.GetAllBooks(ctx, name)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: name, Message: "h.bookUC.GetAllBooks got an error on BookHandler.GetBooks"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	// var booksResponse []map[string]interface{}
	// for _, v := range books {
	// 	formatter := map[string]interface{}{

	// 	}

	// 	booksResponse = append(booksResponse, formatter)
	// }

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetBooks", Code: http.StatusOK, Success: true}, Data: books})
}

func (h BookHandler) DeleteBookyByID(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	err := h.bookUC.DeleteBookByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.bookUC.DeleteBookByID got an error on BookHandler.DeleteBookyByID"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to DeleteBookyByID", Code: http.StatusOK, Success: true})
}
