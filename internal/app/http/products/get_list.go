package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	httpResponses "go_template_project/internal/app/http/responses"
	productsDomain "go_template_project/internal/domain/products"
	"net/http"
	"strconv"
)

type (
	getListCommand interface {
		GetProducts(ctx context.Context, data productsDomain.GetProductsDTO) ([]productsDomain.Product, error)
	}

	GetListHandler struct {
		name           string
		getListCommand getListCommand
	}

	getListRequest struct {
		params productsDomain.GetProductsDTO
	}
)

func NewProductsGetHandler(command getListCommand, name string) *GetListHandler {
	return &GetListHandler{
		name:           name,
		getListCommand: command,
	}
}

// @Summary		Get products
// @Description	Get products list by limit and offset
// @Tags			Products
// @Produce		json
// @Param			limit	query		int						false	"List limit"	default(50)	max(50)
// @Param			offset	query		int						false	"List offset"	default(0)
// @Success		200		{array}		productsDomain.Product	"Product"
// @Failure		400		{string}	string					"Bad Request"
// @Failure		500		{string}	string					"Internal Server Error"
// @Router			/api/products/ [get]
func (h *GetListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *getListRequest
		err         error
	)

	if requestData, err = h.getRequestData(r); err != nil {
		httpResponses.GetResponse(
			w,
			h.name,
			err,
			http.StatusBadRequest,
			nil,
		)
		return
	}

	if err = h.validateRequestData(requestData); err != nil {
		httpResponses.GetResponse(
			w,
			h.name,
			err,
			http.StatusBadRequest,
			nil,
		)
		return
	}

	responseRawBody, err := h.getListCommand.GetProducts(ctx, requestData.params)

	if err != nil {
		httpResponses.GetResponse(
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
		httpResponses.GetResponse(
			w,
			h.name,
			fmt.Errorf("json marshalling failed: %w", err),
			http.StatusInternalServerError,
			nil,
		)
		return
	}

	httpResponses.GetResponse(
		w,
		h.name,
		nil,
		http.StatusOK,
		&responseBody,
	)
}

func (h *GetListHandler) getRequestData(r *http.Request) (requestData *getListRequest, err error) {
	requestData = &getListRequest{}

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

func (h *GetListHandler) validateRequestData(requestData *getListRequest) error {
	return validator.New().Struct(requestData)
}
