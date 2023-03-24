package categorystorage

import "github.com/hvchien216/golang-ecommerce/internal/pkg/db/pg"

type sqlStore struct {
	db pg.PGExecutor
}

func NewSQLStore(db pg.PGExecutor) *sqlStore {
	return &sqlStore{db: db}
}
