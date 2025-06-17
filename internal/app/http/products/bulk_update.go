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
	bulkUpdateCommand interface {
		BulkUpdateProducts(ctx context.Context, data []productsDomain.Product) ([]productsDomain.Product, error)
	}

	BulkUpdateHandler struct {
		name              string
		bulkUpdateCommand bulkUpdateCommand
	}

	bulkUpdateRequest struct {
		body []productsDomain.Product
	}
)

func NewProductBulkUpdateHandler(command bulkUpdateCommand, name string) *BulkUpdateHandler {
	return &BulkUpdateHandler{
		name:              name,
		bulkUpdateCommand: command,
	}
}

// @Summary		Bulk update products
// @Description	Bulk update products
// @Tags			Products
// @Produce		json
// @Success		200	array		productsDomain.Product	"Products"
// @Failure		400	{string}	string					"Bad Request"
// @Failure		500	{string}	string					"Internal Server Error"
// @Router			/api/products [patch]
func (h *BulkUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		ctx         = r.Context()
		requestData *bulkUpdateRequest
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

	responseRawBody, err := h.bulkUpdateCommand.BulkUpdateProducts(ctx, requestData.body)
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

func (h *BulkUpdateHandler) getRequestData(r *http.Request) (requestData *bulkUpdateRequest, err error) {
	requestData = &bulkUpdateRequest{}
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

func (h *BulkUpdateHandler) validateRequestData(requestData *bulkUpdateRequest) error {
	return validator.New().Struct(requestData)
}
