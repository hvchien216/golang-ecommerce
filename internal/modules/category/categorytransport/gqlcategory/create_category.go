package gqlcategory

import (
	"context"
	"github.com/friendsofgo/errors"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorybiz"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (i impl) CreateCategory(ctx context.Context, input categorymodel.NewCategoryInput) (*orm.Category, error) {

	store := categorystorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	biz := categorybiz.New(store)

	cat, err := biz.CreateCategory(ctx, input)

	if err != nil {
		return nil, errors.New("Asd")
	}

	return cat, nil
}
