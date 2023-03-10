package pg

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PGExecutor interface {
	boil.ContextExecutor
	Close() error
}

type pool struct {
	*sql.DB
}

func (p pool) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.DB.ExecContext(ctx, query, args...)
}

func (p pool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.DB.QueryContext(ctx, query, args...)
}

func (p pool) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.DB.QueryRowContext(ctx, query, args...)
}
