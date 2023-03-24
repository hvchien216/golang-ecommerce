package categorystorage

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *sqlStore) Update(ctx context.Context, id int, name string, description *string, parentId *int) (*orm.Category, error) {
	db := s.db

	cat, err := s.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	cat.Name = null.StringFrom(name)

	if description != nil {
		cat.Description = null.StringFromPtr(description)
	}

	if parentId != nil {
		cat.ParentID = null.IntFromPtr(parentId)
	}
	_, err = cat.Update(ctx, db, boil.Infer())

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return cat, nil
}
