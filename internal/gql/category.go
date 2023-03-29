package gql

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (r mutationResolver) CreateCategory(ctx context.Context, input categorymodel.NewCategoryInput) (*orm.Category, error) {
	return r.categoryResolver.CreateCategory(ctx, input)
}

func (r mutationResolver) UpdateCategory(ctx context.Context, id int, name string, description *string, parentID *int) (*orm.Category, error) {
	return r.categoryResolver.UpdateCategory(ctx, id, name, description, parentID)
}

func (r queryResolver) Categories(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) (*categorymodel.CategoryResult, error) {
	return r.categoryResolver.Categories(ctx, filter, paging)
}

func (r queryResolver) Category(ctx context.Context, id int) (*orm.Category, error) {
	return r.categoryResolver.Category(ctx, id)
}

func (r categoryResolver) Name(ctx context.Context, obj *orm.Category) (string, error) {
	return obj.Name, nil
}

func (r categoryResolver) Description(ctx context.Context, obj *orm.Category) (*string, error) {
	return &obj.Description.String, nil
}
func (r categoryResolver) Parent(ctx context.Context, obj *orm.Category) (*orm.Category, error) {
	return r.categoryResolver.Parent(ctx, obj)
}
func (r categoryResolver) Children(ctx context.Context, obj *orm.Category) ([]*orm.Category, error) {
	return r.categoryResolver.Children(ctx, obj)
}
