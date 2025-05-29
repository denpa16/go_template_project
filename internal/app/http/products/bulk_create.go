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
	bulkCreateCommand interface {
		BulkCreateProducts(ctx context.Context, data []productsDomain.Product) (int64, error)
	}

	BulkCreateHandler struct {
		name              string
		bulkCreateCommand bulkCreateCommand
	}

	bulkCreateRequest struct {
		body []productsDomain.Product
	}
)

func NewProductBulkCreateHandler(command bulkCreateCommand, name string) *BulkCreateHandler {
	return &BulkCreateHandler{
		name:              name,
		bulkCreateCommand: command,
	}
}

// @Summary		Bulk create products
// @Description	Bulk create products
// @Tags			Products
// @Produce		json
// @Success		201	{int}		int
// @Failure		400	{string}	string	"Bad Request"
// @Failure		500	{string}	string	"Internal Server Error"
// @Router			/api/products [post]
func (h *BulkCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *bulkCreateRequest
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

	responseRawBody, err := h.bulkCreateCommand.BulkCreateProducts(ctx, requestData.body)
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
		http.StatusCreated,
		&responseBody,
	)
}

func (h *BulkCreateHandler) getRequestData(r *http.Request) (requestData *bulkCreateRequest, err error) {
	requestData = &bulkCreateRequest{}
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
	bodyData := &[]productsDomain.Product{}
	err = json.Unmarshal(body, bodyData)
	if err != nil {
		log.Println(err)
		return
	}
	requestData.body = *bodyData

	return
}

func (h *BulkCreateHandler) validateRequestData(requestData *bulkCreateRequest) error {
	return validator.New().Struct(requestData)
}
