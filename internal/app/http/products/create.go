package products

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	httpResponses "go_template_project/internal/app/http/responses"
	productsDomain "go_template_project/internal/domain/products"
	"io"
	"log"
	"net/http"
)

type (
	createCommand interface {
		CreateProduct(ctx context.Context, data productsDomain.CreateProductDTO) (*productsDomain.Product, error)
	}

	CreateHandler struct {
		name          string
		createCommand createCommand
	}

	createRequest struct {
		body productsDomain.CreateProductDTO
	}
)

func NewProductCreateHandler(command createCommand, name string) *CreateHandler {
	return &CreateHandler{
		name:          name,
		createCommand: command,
	}
}

// @Summary		Create product
// @Description	Create product by id
// @Tags			Products
// @Produce		json
// @Success		201		{object}		domain.Product	"Product"
// @Failure		400		{string}	string			"Bad Request"
// @Failure		500		{string}	string			"Internal Server Error"
// @Router			/api/products [post]
func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *createRequest
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

	responseRawBody, err := h.createCommand.CreateProduct(ctx, requestData.body)

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

func (h *CreateHandler) getRequestData(r *http.Request) (requestData *createRequest, err error) {
	requestData = &createRequest{}
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
	bodyData := &productsDomain.CreateProductDTO{}
	err = json.Unmarshal(body, bodyData)
	if err != nil {
		log.Println(err)
		return
	}
	requestData.body = *bodyData

	return
}

func (h *CreateHandler) validateRequestData(requestData *createRequest) error {
	return validator.New().Struct(requestData)
}
