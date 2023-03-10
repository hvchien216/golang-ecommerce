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
		Description: null.StringFromPtr(&data.Description),
		Slug:        null.StringFromPtr(data.Slug),
		Code:        null.StringFromPtr(data.Code),
		Image:       null.StringFromPtr(data.Image),
		Images:      null.StringFromPtr(data.Images),
		Price:       null.Float64{Float64: data.Price},
		PriceMin:    null.Float64FromPtr(data.PriceMin),
		PriceMax:    null.Float64FromPtr(data.PriceMax),
		Status:      null.String{String: data.Status.String()},
		IsComplete:  null.Bool{Bool: data.IsComplete},
	}

	err := product.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	data.Id = product.ID
	data.CreatedAt = product.CreatedAt
	data.UpdatedAt = product.UpdatedAt.Time

	return data, nil
}
