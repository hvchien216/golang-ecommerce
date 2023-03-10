package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

func NewPool(ctx context.Context, url string) (PGExecutor, error) {
	connCfg, err := pgx.ParseConfig(url)

	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("parsing pgx config failed. err: %w", err))
	}

	connStr := stdlib.RegisterConnConfig(connCfg)
	connInstance, err := sql.Open("pgx", connStr)

	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("parsing pgx config failed. err: %w", err))
	}

	if err = connInstance.PingContext(ctx); err != nil {
		return nil, errors.WithStack(fmt.Errorf("unable to ping DB. err: %w", err))
	}

	return &pool{connInstance}, nil
}
