package gqlcategory

import (
	"context"
	"fmt"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorybiz"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"strconv"
)

func (i impl) Categories(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) (*categorymodel.CategoryResult, error) {

	store := categorystorage.NewSQLStore(i.appCtx.GetMainDBConnection())
	biz := categorybiz.New(store)

	cats, count, err := biz.ListCategory(ctx, filter, paging)
	fmt.Println("biz.ListCategory===>", cats, err, count)

	//pageInfo := categorymodel.PageInfo{
	//	Page:    null.IntFromPtr(paging.Page),
	//	Limit:   paging.Limit,
	//	OrderBy: paging.OrderBy,
	//	Asc:     paging.Asc,
	//	Total:   count,
	//	Cursor:  paging.Cursor,
	//}

	pageInfo := NewPageInfo(paging, count)

	for i := range cats {

		if i == len(cats)-1 {
			pageInfo.NextCursor = strconv.Itoa(cats[i].ID)
		}
	}

	result := categorymodel.CategoryResult{Categories: cats, PageInfo: pageInfo}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func NewPageInfo(paging *common.Paging, total int64) *categorymodel.PageInfo {
	pageInfo := &categorymodel.PageInfo{
		Total: total,
	}

	if paging.Page != nil {
		pageInfo.Page = *paging.Page
	}
	if paging.Limit != nil {
		pageInfo.Limit = *paging.Limit
	}
	if paging.OrderBy != nil {
		pageInfo.OrderBy = *paging.OrderBy
	}
	if paging.Asc != nil {
		pageInfo.Asc = *paging.Asc
	}
	if paging.Cursor != nil {
		pageInfo.Cursor = paging.Cursor
	}
	return pageInfo
}
