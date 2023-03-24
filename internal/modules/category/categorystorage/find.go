package categorystorage

import (
	"context"
	"database/sql"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *sqlStore) FindById(ctx context.Context, id int) (*orm.Category, error) {
	db := s.db

	// Find and return the parent category from the database
	parent, err := orm.Categories(orm.CategoryWhere.ID.EQ(id), qm.Load(orm.CategoryRels.Parent)).One(ctx, db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return parent, nil
}
