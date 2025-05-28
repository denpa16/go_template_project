package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	httpResponses "go_template_project/internal/app/http/responses"
	productsDomain "go_template_project/internal/domain/products"
	"net/http"
)

type (
	getCommand interface {
		GetProduct(ctx context.Context, data productsDomain.GetProductDTO) (*productsDomain.Product, error)
	}

	GetHandler struct {
		name       string
		getCommand getCommand
	}

	getRequest struct {
		params productsDomain.GetProductDTO
	}
)

func NewProductGetHandler(command getCommand, name string) *GetHandler {
	return &GetHandler{
		name:       name,
		getCommand: command,
	}
}

// @Summary		Get product
// @Description	Get product by id
// @Tags			Products
// @Produce		json
// @Success		200		{object}		domain.Product	"Product"
// @Failure		400		{string}	string			"Bad Request"
// @Failure		500		{string}	string			"Internal Server Error"
// @Router			/api/products/{id} [get]
func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *getRequest
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

	responseRawBody, err := h.getCommand.GetProduct(ctx, requestData.params)

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

func (h *GetHandler) getRequestData(r *http.Request) (requestData *getRequest, err error) {
	requestData = &getRequest{}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	requestData.params.ID = id
	return
}

func (h *GetHandler) validateRequestData(requestData *getRequest) error {
	return validator.New().Struct(requestData)
}
