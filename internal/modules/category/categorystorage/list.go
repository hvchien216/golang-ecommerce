package categorystorage

import (
	"context"
	"fmt"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *sqlStore) ListByConditionWithCount(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]*orm.Category, int64, error) {
	db := s.db

	mods := []qm.QueryMod{
		qm.Select("id, name, description, parent_id, created_at, updated_at"),
		qm.Load(orm.CategoryRels.Parent),
	}

	if filter != nil {
		if filter.Id != nil {
			mods = append(mods, orm.CategoryWhere.ParentID.EQ(null.IntFromPtr(filter.Id)))
		}

		if filter.Keyword != nil {
			search := fmt.Sprintf("%%%s%%", *filter.Keyword)
			mods = append(mods,
				qm.Or("name ilike ?", search),
				qm.Or("description ilike ?", search),
			)
		}
	}

	count, err := orm.Categories(mods...).Count(ctx, db)
	if err != nil {
		return orm.CategorySlice{}, 0, err
	}

	if paging.Limit != nil {
		mods = append(mods, qm.Limit(int(*paging.Limit)))
	}

	if v := paging.Cursor; v != nil && *v != "" {
		mods = append(mods, qm.Where("id > ?", paging.Cursor))
	} else {
		var offset int
		if paging.Page != nil && paging.Limit != nil {
			offset = (*paging.Page - 1) * (*paging.Limit)
		}
		mods = append(mods, qm.Offset(offset))
	}

	//countMods := make([]qm.QueryMod, len(mods))
	//copy(countMods, mods)
	//
	//countMods = append(countMods, qm.Offset(0))

	if paging.OrderBy != nil && paging.Asc != nil {
		var order string
		if paging.Asc != nil && *paging.Asc == false {
			order = "desc"
		} else {
			order = "asc"

		}

		mods = append([]qm.QueryMod{qm.OrderBy(fmt.Sprintf("%s %s", *paging.OrderBy, order))}, mods...)
	}
	cats, err := orm.Categories(mods...).All(ctx, db)

	if err != nil {
		return orm.CategorySlice{}, 0, err
	}

	//mods = append(mods, qm.Offset(0))

	return cats, count, nil
}

//&gqlerror.Error{
//Message: "not authorized",
//Extensions: map[string]interface{}{
//"code": "401",
//},
//}
