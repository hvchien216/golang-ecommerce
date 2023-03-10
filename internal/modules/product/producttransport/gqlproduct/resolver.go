package gqlproduct

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
	"github.com/hvchien216/golang-ecommerce/internal/pkg/component/appctx"
)

type Resolver interface {
	CreateProduct(context.Context, productmodel.NewProductInput) (*productmodel.Product, error)
}

func New(appCtx appctx.AppContext) Resolver {
	return impl{appCtx: appCtx}
}

type impl struct {
	appCtx appctx.AppContext
}
