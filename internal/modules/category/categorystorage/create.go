package categorystorage

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (s *sqlStore) Create(ctx context.Context, input categorymodel.NewCategoryInput) (*orm.Category, error) {
	db := s.db

	category := &orm.Category{
		Name:        null.StringFrom(input.Name),
		Description: null.StringFrom(input.Description),
		ParentID:    null.IntFromPtr(input.ParentID),
	}

	err := category.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	//if category.ParentID.Valid {
	//	parentModel, err := orm.Categories(
	//		qm.Where("id = ?", category.ParentID),
	//	).One(ctx, db)
	//	if err != nil {
	//		return nil, err
	//	}
	//	parent := &categorymodel.Category{
	//		ID:          parentModel.ID,
	//		Name:        parentModel.Name.String,
	//		Description: parentModel.Description.String,
	//		CreatedAt:   &parentModel.CreatedAt,
	//		UpdatedAt:   &parentModel.UpdatedAt,
	//	}
	//	category.Parent = parent
	//}
	//if input.ParentID != nil {
	//	parentCategory, err := models.Categories(qm.Where("id = ?", *input.ParentID)).One(ctx, db)
	//	if err != nil {
	//		return nil, err
	//	}
	//	newCategory.ParentID = parentCategory.ID
	//	newCategory.Parent = &Category{
	//		ID:          parentCategory.ID,
	//		Name:        parentCategory.Name,
	//		Description: parentCategory.Description,
	//		CreatedAt:   parentCategory.CreatedAt,
	//		UpdatedAt:   parentCategory.UpdatedAt,
	//	}
	//}

	return category, nil
}

//func mapORMToModel(ormCategory *orm.Category) *categorymodel.Category {
//	if ormCategory == nil {
//		return nil
//	}
//
//	modelCategory := &categorymodel.Category{
//		ID:          ormCategory.ID,
//		Name:        ormCategory.Name.String,
//		Description: ormCategory.Description.String,
//		CreatedAt:   &ormCategory.CreatedAt,
//		UpdatedAt:   &ormCategory.UpdatedAt,
//	}
//
//	if ormCategory.R.Parent. {
//		parentCategory := mapORMToModel(ormCategory.R.Parent.Category)
//		modelCategory.Parent = parentCategory
//	}
//
//	if ormCategory.R.Children != nil {
//		for _, child := range ormCategory.R.Children {
//			childCategory := mapORMToModel(child)
//			modelCategory.Children = append(modelCategory.Children, childCategory)
//		}
//	}
//
//	return modelCategory
//}
