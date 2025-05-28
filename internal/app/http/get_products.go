package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go_template_project/internal/domain"
	"net/http"
	"strconv"
)

type (
	getProductsGetCommand interface {
		GetProducts(ctx context.Context, data domain.GetProductsDTO) ([]domain.Product, error)
	}

	ProductsGetHandler struct {
		name                  string
		getProductsGetCommand getProductsGetCommand
	}

	getProductsGetRequest struct {
		params domain.GetProductsDTO
	}
)

func NewProductsGetHandler(command getProductsGetCommand, name string) *ProductsGetHandler {
	return &ProductsGetHandler{
		name:                  name,
		getProductsGetCommand: command,
	}
}

// @Summary		Get products
// @Description	Get products list by limit and offset
// @Tags			Products
// @Produce		json
// @Param			limit	query		int				false	"List limit"	default(50)	max(50)
// @Param			offset	query		int				false	"List offset"
// @Success		200		{array}		domain.Product	"Product"
// @Failure		400		{string}	string			"Bad Request"
// @Failure		500		{string}	string			"Internal Server Error"
// @Router			/api/products/ [get]
func (h *ProductsGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *getProductsGetRequest
		err         error
	)

	if requestData, err = h.getRequestData(r); err != nil {
		GetResponse(
			w,
			h.name,
			err,
			http.StatusBadRequest,
			nil,
		)
		return
	}

	if err = h.validateRequestData(requestData); err != nil {
		GetResponse(
			w,
			h.name,
			err,
			http.StatusBadRequest,
			nil,
		)
		return
	}

	responseRawBody, err := h.getProductsGetCommand.GetProducts(ctx, requestData.params)

	if err != nil {
		GetResponse(
			w,
			h.name,
			fmt.Errorf("command handler failed: %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}

	responseBody, err := json.Marshal(responseRawBody)
	if err != nil {
		GetResponse(
			w,
			h.name,
			fmt.Errorf("json marshalling failed: %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}

	GetResponse(
		w,
		h.name,
		nil,
		http.StatusOK,
		&responseBody,
	)
}

func (h *ProductsGetHandler) getRequestData(r *http.Request) (requestData *getProductsGetRequest, err error) {
	requestData = &getProductsGetRequest{}

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		limit = 50
	}
	if limit == 0 || limit > 50 {
		limit = 50
	}
	requestData.params.Limit = int64(limit)

	offset, err := strconv.Atoi(r.FormValue("offset"))
	if err != nil {
		offset = 0
	}
	requestData.params.Offset = int64(offset)

	name := r.FormValue("name")
	if name != "" {
		requestData.params.Name = name
	}
	title := r.FormValue("title")
	if title != "" {
		requestData.params.Title = title
	}

	return
}

func (h *ProductsGetHandler) validateRequestData(requestData *getProductsGetRequest) error {
	return validator.New().Struct(requestData)
}
