package categorybiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/common"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
)

func (biz business) FindCategory(
	ctx context.Context,
	id int) (*orm.Category, error) {

	//data.Status = productmodel.ProductStatusOutOfStock
	//data.IsComplete = false

	//if err := data.Validate(); err != nil {
	//	return err
	//}

	cat, err := biz.store.FindById(ctx, id)

	if err != nil {
		if err != common.RecordNotFound {
			//return nil, common.ErrEntityNotFound(categorymodel.EntityName, err)
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)

		}
		// OR: able to throw err `sth went wrong with server`
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	return cat, nil
}
