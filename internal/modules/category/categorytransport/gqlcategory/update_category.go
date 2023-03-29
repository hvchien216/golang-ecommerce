package gqlcategory

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorybiz"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (i impl) UpdateCategory(
	ctx context.Context,
	id int,
	name string,
	description *string,
	parentID *int,
) (*orm.Category, error) {

	store := categorystorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	biz := categorybiz.New(store)

	cat, err := biz.UpdateCategory(ctx, id, name, description, parentID)

	if err != nil {
		return nil, err
	}

	return cat, nil
}
