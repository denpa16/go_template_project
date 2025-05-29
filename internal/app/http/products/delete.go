package products

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	httpResponses "go_template_project/internal/app/http/responses"
	productsDomain "go_template_project/internal/domain/products"
	"net/http"
)

type (
	deleteCommand interface {
		DeleteProduct(ctx context.Context, data productsDomain.DeleteProductDTO) (*productsDomain.Product, error)
	}

	DeleteHandler struct {
		name          string
		deleteCommand deleteCommand
	}

	deleteRequest struct {
		params productsDomain.DeleteProductDTO
	}
)

func NewProductDeleteHandler(command deleteCommand, name string) *DeleteHandler {
	return &DeleteHandler{
		name:          name,
		deleteCommand: command,
	}
}

// @Summary		Delete product
// @Description	Delete product by id
// @Tags			Products
// @Produce		json
// @Success		204	{object}	string	"No content"
// @Failure		400	{string}	string	"Bad Request"
// @Failure		404	{string}	string	"Not Found"
// @Failure		500	{string}	string	"Internal Server Error"
// @Router			/api/products/{id} [delete]
func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *deleteRequest
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

	responseRawBody, err := h.deleteCommand.DeleteProduct(ctx, requestData.params)

	if err != nil {
		if errors.Is(err, productsDomain.ErrProductNotFound) {
			httpResponses.GetResponse(
				w,
				h.name,
				err,
				http.StatusNotFound,
				nil,
			)
			return
		}
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
		http.StatusNoContent,
		&responseBody,
	)
}

func (h *DeleteHandler) getRequestData(r *http.Request) (requestData *deleteRequest, err error) {
	requestData = &deleteRequest{}
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	requestData.params.ID = id
	return
}

func (h *DeleteHandler) validateRequestData(requestData *deleteRequest) error {
	return validator.New().Struct(requestData)
}
