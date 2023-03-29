package gqlcategory

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorybiz"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (i impl) Category(ctx context.Context, id int) (*orm.Category, error) {

	store := categorystorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	biz := categorybiz.New(store)

	cat, err := biz.FindCategory(ctx, id)

	if err != nil {
		return nil, err
	}

	return cat, nil
}
