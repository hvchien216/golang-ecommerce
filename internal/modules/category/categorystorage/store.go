package categorystorage

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/hvchien216/golang-ecommerce/internal/pkg/db/pg"
)

type Store interface {
	ListByConditionWithCount(context.Context, *categorymodel.Filter, *common.Paging) ([]*orm.Category, int64, error)

	FindById(context.Context, int) (*orm.Category, error)

	Create(context.Context, categorymodel.NewCategoryInput) (*orm.Category, error)

	Update(ctx context.Context, id int, name string, description *string, parentID *int) (*orm.Category, error)
}

type sqlStore struct {
	db pg.PGExecutor
}

func NewSQLStore(db pg.PGExecutor) *sqlStore {
	return &sqlStore{db: db}
}
