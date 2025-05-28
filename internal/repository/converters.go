package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func NConvertPgTimestamp(value pgtype.Timestamp) *time.Time {
	if value.Valid {
		return &value.Time
	}
	return nil
}
