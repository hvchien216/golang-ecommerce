package gqlproduct

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productbiz"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productstorage"
)

func (i impl) CreateProduct(ctx context.Context, input productmodel.NewProductInput) (*productmodel.Product, error) {

	data := productmodel.Product{
		Name:        input.Name,
		Description: input.Description,
		Slug:        input.Slug,
		Code:        input.Code,
		Image:       input.Image,
		Images:      input.Images,
		Price:       input.Price,
		PriceMin:    input.PriceMin,
		PriceMax:    input.PriceMax,
	}

	store := productstorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	biz := productbiz.NewCreateProductBiz(store)

	product, err := biz.CreateProduct(ctx, &data)

	if err != nil {
		return nil, err
	}

	return product, nil
}
