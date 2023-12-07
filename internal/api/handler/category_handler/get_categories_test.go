package category_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/category_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	type resp struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	testCases := []struct {
		name       string
		categories repository.Categories
		mockError  error
		expected   []resp
		wantErr    error
	}{
		{
			name: "Success",
			categories: repository.Categories{
				{ID: 1, Name: "Category 1"},
				{ID: 2, Name: "Category 2"},
			},
			mockError: nil,
			expected: []resp{
				{ID: 1, Name: "Category 1"},
				{ID: 2, Name: "Category 2"},
			},
		},
		{
			name:       "Error",
			categories: nil,
			mockError:  tests.ErrTest,
			expected:   nil,
			wantErr:    tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("GetAll", ctx).Return(tc.categories, tc.mockError)

			cs := service.NewCategoryService(mockRepo)

			response := category_handler.GetCategories(ctx, cs)(nil, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
