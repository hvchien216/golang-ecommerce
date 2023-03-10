package gql

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
)

func (m mutationResolver) CreateProduct(ctx context.Context, input productmodel.NewProductInput) (*productmodel.Product, error) {
	return m.productResolver.CreateProduct(ctx, input)
}
