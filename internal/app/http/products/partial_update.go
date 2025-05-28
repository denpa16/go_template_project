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
	"io"
	"log"
	"net/http"
)

type (
	partialUpdateCommand interface {
		PartialUpdateProduct(
			ctx context.Context,
			data productsDomain.PartialUpdateProductDTO,
		) (*productsDomain.Product, error)
	}

	PartialUpdateHandler struct {
		name                 string
		partialUpdateCommand partialUpdateCommand
	}

	partialUpdateRequest struct {
		data productsDomain.PartialUpdateProductDTO
	}
)

func NewProductPartialUpdateHandler(command partialUpdateCommand, name string) *PartialUpdateHandler {
	return &PartialUpdateHandler{
		name:                 name,
		partialUpdateCommand: command,
	}
}

// @Summary		PartialUpdate product
// @Description	PartialUpdate product by id
// @Tags			Products
// @Produce		json
// @Success		200		{object}		domain.Product	"Product"
// @Failure		400		{string}	string			"Bad Request"
// @Failure		404		{string}	string			"Not found"
// @Failure		500		{string}	string			"Internal Server Error"
// @Router			/api/products/{id} [patch]
func (h *PartialUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *partialUpdateRequest
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

	responseRawBody, err := h.partialUpdateCommand.PartialUpdateProduct(ctx, requestData.data)

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
		http.StatusOK,
		&responseBody,
	)
}

func (h *PartialUpdateHandler) getRequestData(r *http.Request) (requestData *partialUpdateRequest, err error) {
	requestData = &partialUpdateRequest{}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)
	bodyData := &productsDomain.PartialUpdateProductDTO{}
	err = json.Unmarshal(body, bodyData)
	if err != nil {
		log.Println(err)
		return
	}
	requestData.data = *bodyData

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	requestData.data.ID = id

	return
}

func (h *PartialUpdateHandler) validateRequestData(requestData *partialUpdateRequest) error {
	return validator.New().Struct(requestData)
}
