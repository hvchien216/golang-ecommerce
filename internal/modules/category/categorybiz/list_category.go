package categorybiz

import (
	"context"
	"fmt"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (biz business) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]*orm.Category, int64, error) {

	cats, total, err := biz.store.ListByConditionWithCount(ctx, filter, paging)

	fmt.Println("cats===>", cats, err, total)

	if err != nil {
		return nil, 0, err
	}

	return cats, total, nil
}
