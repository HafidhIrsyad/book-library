package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	api "github.com/book-library/app/helper"
	u "github.com/book-library/app/usecase"
	"github.com/book-library/entity/author"
	"github.com/book-library/logger"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AuthorHandler struct {
	authorUC u.AuthorServiceI
}

func NewAuthorHandler(authorUC u.AuthorServiceI) AuthorHandler {
	return AuthorHandler{
		authorUC: authorUC,
	}
}

func (h AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())

	var input author.AuthorInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on AuthorHandler.CreateAuthor"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.authorUC.CreateAuthor(ctx, input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.authorUC.CreateAuthor got an error on AuthorHandler.CreateAuthor"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to CreateAuthor", Code: http.StatusOK, Success: true})
}

func (h AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	var input author.AuthorInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on AuthorHandler.UpdateAuthor"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.authorUC.UpdateAuthor(ctx, int64(idInt), input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.categoryUC.UpdateAuthor got an error on AuthorHandler.UpdateAuthor"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to UpdateAuthor", Code: http.StatusOK, Success: true})
}

func (h AuthorHandler) GetAuhtorById(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	catById, err := h.authorUC.GetAuthorByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.categoryUC.GetCategoryById got an error on AuthorHandler.GetAuhtorById"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetAuhtorById", Code: http.StatusOK, Success: true}, Data: catById})
}

func (h AuthorHandler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	name := r.URL.Query().Get("name")

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: name})

	categories, err := h.authorUC.GetAllAuthors(ctx, name)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: name, Message: "h.authorUC.GetAllAuthors got an error on AuthorHandler.GetAuthors"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetAuthors", Code: http.StatusOK, Success: true}, Data: categories})
}

func (h AuthorHandler) DeleteAuthorByID(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	err := h.authorUC.DeleteAuthorByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.authorUC.DeleteAuthorByID got an error on AuthorHandler.DeleteAuthorByID"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to DeleteAuthorByID", Code: http.StatusOK, Success: true})
}
