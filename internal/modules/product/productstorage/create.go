package productstorage

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/product/productmodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.Product) (*productmodel.Product, error) {
	db := s.db

	product := orm.Product{
		Name:        data.Name,
		Description: null.StringFrom(data.Description),
		Slug:        null.StringFrom(data.Slug),
		Code:        null.StringFrom(data.Code),
		Image:       null.StringFrom(data.Image),
		Images:      null.StringFrom(data.Images),
		Price:       null.Float64From(data.Price),
		PriceMin:    null.Float64From(data.PriceMin),
		PriceMax:    null.Float64From(data.PriceMax),
		Status:      null.StringFrom(data.Status.String()),
		IsComplete:  null.BoolFrom(data.IsComplete),
	}

	err := product.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	data.Id = product.ID
	data.CreatedAt = product.CreatedAt
	data.UpdatedAt = product.UpdatedAt

	return data, nil
}
