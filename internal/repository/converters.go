package repository

import (
	"encoding/json"
	"github.com/jackc/pgx/v5/pgtype"
	"go_template_project/internal/domain"
	"time"
)

func NConvertPgTimestamp(value pgtype.Timestamp) *time.Time {
	if value.Valid {
		return &value.Time
	}

	return nil
}

func getProductsParamsRaw(params domain.GetProductsDTO) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
