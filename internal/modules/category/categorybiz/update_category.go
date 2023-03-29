package categorybiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (biz business) UpdateCategory(
	ctx context.Context,
	id int,
	name string,
	description *string,
	parentID *int,
) (*orm.Category, error) {

	//data.Status = productmodel.ProductStatusOutOfStock
	//data.IsComplete = false

	//if err := data.Validate(); err != nil {
	//	return err
	//}

	cat, err := biz.store.Update(ctx, id, name, description, parentID)

	if err != nil {
		return nil, err
	}

	return cat, nil
}
