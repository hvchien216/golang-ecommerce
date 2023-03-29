package gqlcategory

import (
	"testing"
)

func TestImpl_CreateCategory(t *testing.T) {
	//tests := map[string]struct {
	//	input        categorymodel.NewCategoryInput
	//	categoryCtrl *struct {
	//		input  categorymodel.NewCategoryInput
	//		output *orm.Category
	//		err    error
	//	}
	//	wantErr error
	//}{
	//	"success": {
	//		input: categorymodel.NewCategoryInput{
	//			Name:        "cat 01",
	//			Description: "cat desc 01",
	//		},
	//		categoryCtrl: &struct {
	//			input  categorymodel.NewCategoryInput
	//			output *orm.Category
	//			err    error
	//		}{input: categorymodel.NewCategoryInput{
	//			Name:        "cat 01",
	//			Description: "cat desc 01",
	//		}, output: &orm.Category{
	//			ID:   1,
	//			Name: "cat 01",
	//			Description: null.String{
	//				String: "cat desc 01",
	//				Valid:  true,
	//			},
	//			CreatedAt: time.Date(2022, 8, 7, 9, 10, 11, 12, time.UTC),
	//			UpdatedAt: time.Date(2022, 8, 7, 9, 10, 11, 12, time.UTC),
	//		}},
	//	},
	//}

	//for desc, tt := range tests {
	//	t.Run(desc, func(t *testing.T) {
	//		ctx := context.Background()
	//
	//		categoryStore := new(categorybiz.MockCreateCategoryStore)
	//		if tt.categoryCtrl != nil {
	//			categoryStore.ExpectedCalls = []*mock.Call{
	//				categoryStore.On("Create", ctx, tt.categoryCtrl.input).Return(tt.categoryCtrl.output, tt.categoryCtrl.err),
	//			}
	//		}
	//
	//		m := &mutationResolver{&resolver{
	//			batchesCtrl: batchCtrl,
	//		}}
	//
	//		rs, err := m.RequestBatchApproval(ctx, tt.givenBatchExternalID)
	//		if tt.wantErr != nil {
	//			require.EqualError(t, tt.wantErr, err.Error())
	//		} else {
	//			require.NoError(t, err)
	//			require.Equal(t, mod.NewBatch(tt.batchCtrl.output), rs)
	//		}
	//	})
	//}
}
