package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	api "github.com/book-library/app/helper"
	u "github.com/book-library/app/usecase"
	"github.com/book-library/entity/category"
	"github.com/book-library/logger"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CategoryHandler struct {
	categoryUC u.CategoryServiceI
}

func NewCategoryHandler(categoryUC u.CategoryServiceI) CategoryHandler {
	return CategoryHandler{
		categoryUC: categoryUC,
	}
}

func (h CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())

	var input category.CategoryInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on CategoryHandler.CreateCategory"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.categoryUC.CreateCategory(ctx, input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.categoryUC.CreateCategory got an error on CategoryHandler.CreateCategory"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to CreateCategory", Code: http.StatusOK, Success: true})
}

func (h CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	var input category.CategoryInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Message: "json.NewDecoder got an error on CategoryHandler.UpdateCategory"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusInternalServerError, Success: false})
		return
	}

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: input})

	err = h.categoryUC.UpdateCategory(ctx, int64(idInt), input)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: input, Message: "h.categoryUC.UpdateCategory got an error on CategoryHandler.UpdateCategory"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to UpdateCategory", Code: http.StatusOK, Success: true})
}

func (h CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	catById, err := h.categoryUC.GetCategoryByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.categoryUC.GetCategoryById got an error on CategoryHandler.GetCategoryById"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetCategoryById", Code: http.StatusOK, Success: true}, Data: catById})
}

func (h CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	name := r.URL.Query().Get("name")

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: name})

	categories, err := h.categoryUC.GetAllCategories(ctx, name)
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: name, Message: "h.categoryUC.GetAllCategories got an error on CategoryHandler.GetCategories"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponse(w, api.Response{Meta: api.Meta{Message: "Success to GetCategories", Code: http.StatusOK, Success: true}, Data: categories})
}

func (h CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	log := log.With().Str("request_id", uuid.New().String()).Logger()
	ctx := log.WithContext(r.Context())
	id := r.PathValue("id")
	idInt, _ := strconv.Atoi(id)

	logger.LogInfo(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Request: idInt})

	err := h.categoryUC.DeleteCategoryByID(ctx, int64(idInt))
	if err != nil {
		logger.LogError(logger.LogConfig{Logger: log, Req: r, Ctx: ctx, Err: err, Request: idInt, Message: "h.categoryUC.DeleteCategoryByID got an error on CategoryHandler.DeleteCategoryByID"})
		api.APIResponseFailed(w, api.Meta{Message: err.Error(), Code: http.StatusBadRequest, Success: false})
		return
	}

	api.APIResponseSuccessWithoutData(w, api.Meta{Message: "Success to DeleteCategoryByID", Code: http.StatusOK, Success: true})
}
