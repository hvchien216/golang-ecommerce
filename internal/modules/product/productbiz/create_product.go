package productbiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.Product) (*productmodel.Product, error)
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(
	ctx context.Context,
	data *productmodel.Product) (*productmodel.Product, error) {

	data.Status = productmodel.ProductStatusOutOfStock
	data.IsComplete = false

	//if err := data.Validate(); err != nil {
	//	return err
	//}

	product, err := biz.store.Create(ctx, data)

	if err != nil {
		return nil, err
	}

	return product, nil
}
