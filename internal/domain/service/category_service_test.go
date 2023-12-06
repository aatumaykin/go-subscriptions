package service_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_CreateCategory(t *testing.T) {
	testCases := []struct {
		name       string
		category   entity.Category
		wantErr    error
		wantResult *entity.Category
	}{
		{
			name:       "Test empty category name",
			category:   entity.Category{Name: ""},
			wantErr:    service.ErrInvalidCategory,
			wantResult: nil,
		},
		{
			name:       "Test valid category",
			category:   entity.Category{Name: "Test Category"},
			wantErr:    nil,
			wantResult: &entity.Category{ID: 1, Name: "Test Category"},
		},
		{
			name:       "Test error",
			category:   entity.Category{Name: "Test Category"},
			wantErr:    tests.ErrTest,
			wantResult: nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("Create", ctx, tc.category).Return(tc.wantResult, tc.wantErr)

			categoryService := service.NewCategoryService(mockRepo)
			result, err := categoryService.CreateCategory(ctx, tc.category)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCategoryService_GetCategory(t *testing.T) {
	testCases := []struct {
		name       string
		id         uint
		wantErr    error
		wantResult *entity.Category
	}{
		{
			name:       "ID is zero",
			id:         0,
			wantErr:    repository.ErrNotFoundCategory,
			wantResult: nil,
		},
		{
			name:       "Test valid category ID",
			id:         1,
			wantErr:    nil,
			wantResult: &entity.Category{ID: 1, Name: "Test Category"},
		},
		{
			name:       "Test error",
			id:         1,
			wantErr:    tests.ErrTest,
			wantResult: nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("Get", ctx, tc.id).Return(tc.wantResult, tc.wantErr)

			categoryService := service.NewCategoryService(mockRepo)
			result, err := categoryService.GetCategory(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCategoryService_GetAllCategories(t *testing.T) {
	testCases := []struct {
		name       string
		wantErr    error
		wantResult repository.Categories
	}{
		{
			name:       "Test empty categories",
			wantErr:    nil,
			wantResult: repository.Categories{},
		},
		{
			name:       "Test valid categories",
			wantErr:    nil,
			wantResult: repository.Categories{{ID: 1, Name: "Test Category"}, {ID: 2, Name: "Test Category 2"}},
		},
		{
			name:       "Test error",
			wantErr:    tests.ErrTest,
			wantResult: nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("GetAll", ctx).Return(tc.wantResult, tc.wantErr)

			categoryService := service.NewCategoryService(mockRepo)
			result, err := categoryService.GetAllCategories(ctx)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCategoryService_UpdateCategory(t *testing.T) {
	testCases := []struct {
		name       string
		category   entity.Category
		wantErr    error
		wantResult *entity.Category
	}{
		{
			name:       "ID is zero",
			category:   entity.Category{ID: 0, Name: "Test Category"},
			wantErr:    service.ErrInvalidCategory,
			wantResult: nil,
		},
		{
			name:       "Name is an empty string",
			category:   entity.Category{ID: 1, Name: ""},
			wantErr:    service.ErrInvalidCategory,
			wantResult: nil,
		},
		{
			name:       "Test valid category",
			category:   entity.Category{ID: 1, Name: "Test Category"},
			wantErr:    nil,
			wantResult: &entity.Category{ID: 1, Name: "Test Category"},
		},
		{
			name:       "Test error",
			category:   entity.Category{ID: 1, Name: "Test Category"},
			wantErr:    tests.ErrTest,
			wantResult: nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("Update", ctx, tc.category).Return(tc.wantResult, tc.wantErr)

			categoryService := service.NewCategoryService(mockRepo)
			result, err := categoryService.UpdateCategory(ctx, tc.category)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCategoryService_DeleteCategory(t *testing.T) {
	testCases := []struct {
		name    string
		id      uint
		wantErr error
	}{
		{
			name:    "ID is zero",
			id:      0,
			wantErr: repository.ErrNotFoundCategory,
		},
		{
			name:    "Test valid category",
			id:      1,
			wantErr: nil,
		},
		{
			name:    "Test error",
			id:      1,
			wantErr: tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCategoryRepository)
			mockRepo.On("Delete", ctx, tc.id).Return(tc.wantErr)

			categoryService := service.NewCategoryService(mockRepo)
			err := categoryService.DeleteCategory(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}
