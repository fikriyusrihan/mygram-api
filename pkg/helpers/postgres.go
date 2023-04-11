package helpers

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsForeignKeyViolation(err error) bool {
	var pErr *pgconn.PgError
	if errors.As(err, &pErr) {
		return pErr.Code == "23503"
	}
	return false
}
