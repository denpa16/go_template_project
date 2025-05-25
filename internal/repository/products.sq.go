package repository

import (
	"context"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqGetProductsRow struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
}

type SqGetProductsParams struct {
	Limit  int64
	Offset int64
}

func (q *Queries) SqGetProducts(ctx context.Context, arg SqGetProductsParams) ([]SqGetProductsRow, error) {
	preparedParams, err := convertGetProductsParams(arg)
	if err != nil {
		return nil, err
	}
	query := sq.Select(
		"id",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	).From("products")
	if limit, exists := preparedParams["limit"]; exists {
		if limitValue, ok := limit.(int64); ok {
			query.Limit(uint64(limitValue))
		}
	}
	if offset, exists := preparedParams["offset"]; exists {
		if offsetValue, ok := offset.(int64); ok {
			query.Offset(uint64(offsetValue))
		}
	}
	sqlString, args, err := query.ToSql()
	fmt.Println(sqlString)
	if err != nil {
		return nil, fmt.Errorf("sq get products query to sql error: %w", err)
	}
	rows, err := q.db.Query(ctx, sqlString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SqGetProductsRow
	for rows.Next() {
		var i SqGetProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func convertGetProductsParams(params SqGetProductsParams) (map[string]interface{}, error) {
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
