package category_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetCategory(t *testing.T) {
	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name     string
		id       string
		category entity.Category
		expected resp
		wantErr  error
	}{
		{
			name:     "success",
			id:       "1",
			category: entity.Category{ID: 1, Name: "Test Category"},
			expected: resp{
				ID:   1,
				Name: "Test Category",
			},
		},
		{
			name:     "error",
			id:       "10",
			category: entity.Category{ID: 2, Name: "Test Category"},
			wantErr:  repository.ErrNotFoundCategory,
		},
	}

	cs := service.NewCategoryService(memory.NewCategoryRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCategory(ctx, tc.category)
			ps := httprouter.Params{{Key: "id", Value: tc.id}}

			response := category_handler.GetCategory(ctx, cs)(nil, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
