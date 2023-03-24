package gqlcategory

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/hvchien216/golang-ecommerce/internal/pkg/component/appctx"
)

type Resolver interface {
	CreateCategory(context.Context, categorymodel.NewCategoryInput) (*orm.Category, error)
	UpdateCategory(ctx context.Context, id int, name string, description *string, parentID *int) (*orm.Category, error)

	Categories(context.Context, *categorymodel.Filter, *common.Paging) (*categorymodel.CategoryResult, error)
	Category(ctx context.Context, id int) (*orm.Category, error)

	Parent(ctx context.Context, obj *orm.Category) (*orm.Category, error)
	Children(ctx context.Context, obj *orm.Category) ([]*orm.Category, error)
}

func New(appCtx appctx.AppContext) Resolver {
	return impl{appCtx: appCtx}
}

type impl struct {
	appCtx appctx.AppContext
}

func (i impl) Parent(ctx context.Context, obj *orm.Category) (*orm.Category, error) {
	return obj.R.GetParent(), nil
}

func (i impl) Children(ctx context.Context, obj *orm.Category) ([]*orm.Category, error) {
	//store := categorystorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	//biz := categorybiz.NewListCategoryBiz(store)
	//
	//filter := categorymodel.Filter{
	//	Id: &obj.ParentID.Int,
	//}
	//
	//cats, _, err := biz.ListCategory(ctx, &filter, nil)
	//
	//return cats, err
	//children, err := obj.L.

	return nil, nil

}
