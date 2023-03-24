package categorybiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

type CreateCategoryStore interface {
	Create(context.Context, categorymodel.NewCategoryInput) (*orm.Category, error)
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(
	ctx context.Context,
	input categorymodel.NewCategoryInput) (*orm.Category, error) {

	//data.Status = productmodel.ProductStatusOutOfStock
	//data.IsComplete = false

	//if err := data.Validate(); err != nil {
	//	return err
	//}

	cat, err := biz.store.Create(ctx, input)

	if err != nil {
		return nil, err
	}

	return cat, nil
}
