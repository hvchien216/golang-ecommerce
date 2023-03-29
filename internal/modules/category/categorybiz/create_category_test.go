package categorybiz

import (
	"context"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorymodel"
	"github.com/hvchien216/golang-ecommerce/internal/modules/category/categorystorage"
	"github.com/hvchien216/golang-ecommerce/internal/orm"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"testing"
	"time"
)

func TestBusiness_CreateCategory(t *testing.T) {
	dt := time.Date(2022, 10, 27, 0, 0, 0, 0, time.UTC)

	type createCategoryStorage struct {
		input  categorymodel.NewCategoryInput
		output *orm.Category
		err    error
	}

	type mockRepoArg struct {
		createCategoryStorage createCategoryStorage
	}

	tcs := map[string]struct {
		mockRepo         mockRepoArg
		givenCreateInput categorymodel.NewCategoryInput
		want             *orm.Category
		wantErr          error
	}{
		"success": {
			mockRepo: mockRepoArg{
				createCategoryStorage: createCategoryStorage{
					input: categorymodel.NewCategoryInput{
						Name:        "ĐIỆN TỬ GIA DỤNG",
						Description: "Mô tả này nói chung là okela",
					},
					output: &orm.Category{
						ID:          1,
						Name:        "ĐIỆN TỬ GIA DỤNG",
						Description: null.StringFrom("Mô tả này nói chung là okela"),
						CreatedAt:   dt,
						UpdatedAt:   dt,
					},
				},
			},
			givenCreateInput: categorymodel.NewCategoryInput{
				Name:        "ĐIỆN TỬ GIA DỤNG",
				Description: "Mô tả này nói chung là okela",
			},
			want: &orm.Category{
				ID:          1,
				Name:        "ĐIỆN TỬ GIA DỤNG",
				Description: null.StringFrom("Mô tả này nói chung là okela"),
				CreatedAt:   dt,
				UpdatedAt:   dt,
			},
		},
	}

	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			ctx := context.Background()

			mockStorage := categorystorage.MockStore{}
			mockStorage.ExpectedCalls = append(mockStorage.ExpectedCalls, mockStorage.On(
				"Create", ctx, tc.mockRepo.createCategoryStorage.input).Return(
				tc.mockRepo.createCategoryStorage.output,
				tc.mockRepo.createCategoryStorage.err,
			))

			rs, err := New(&mockStorage).CreateCategory(ctx, tc.givenCreateInput)
			if tc.wantErr != nil {
				require.EqualError(t, tc.wantErr, err.Error())
			} else {
				require.Equal(t, tc.want, rs)
			}
		})
	}
}
