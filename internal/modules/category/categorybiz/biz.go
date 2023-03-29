package categorybiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

type Business interface {
	CreateCategory(
		ctx context.Context,
		input categorymodel.NewCategoryInput) (*orm.Category, error)

	FindCategory(
		ctx context.Context,
		id int) (*orm.Category, error)

	ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]*orm.Category, int64, error)

	UpdateCategory(
		ctx context.Context,
		id int,
		name string,
		description *string,
		parentID *int,
	) (*orm.Category, error)
}

type business struct {
	store categorystorage.Store
}

func New(store categorystorage.Store) Business {
	return business{
		store: store,
	}
}
