package memory_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_Create(t *testing.T) {
	testCases := []struct {
		name       string
		category   entity.Category
		wantResult *entity.Category
		wantErr    error
	}{
		{
			name:       "Create a new category",
			category:   entity.Category{Name: "Test Category"},
			wantResult: &entity.Category{ID: 1, Name: "Test Category"},
			wantErr:    nil,
		},
		{
			name:       "Create a new category",
			category:   entity.Category{Name: "Test Category 2"},
			wantResult: &entity.Category{ID: 2, Name: "Test Category 2"},
			wantErr:    nil,
		},
	}

	repo := memory.NewCategoryRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Create(ctx, tc.category)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
			}
		})
	}
}

func TestCategoryRepository_Get(t *testing.T) {
	testCases := []struct {
		name     string
		category entity.Category
		id       uint
		wantErr  error
	}{
		{
			name:     "Get an existing category",
			category: entity.Category{Name: "Test Category"},
			id:       1,
			wantErr:  nil,
		},
		{
			name:     "Get an existing category",
			category: entity.Category{Name: "Test Category 2"},
			id:       2,
			wantErr:  nil,
		},
		{
			name:     "Get a non-existing category",
			category: entity.Category{Name: "Test Category"},
			id:       10,
			wantErr:  repository.ErrNotFoundCategory,
		},
	}

	repo := memory.NewCategoryRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createdCategory, err := repo.Create(ctx, tc.category)
			assert.NoError(t, err)

			foundCategory, err := repo.Get(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, createdCategory, foundCategory)
			}
		})
	}
}

func TestCategoryRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		categories  repository.Categories
		wantResult  repository.Categories
		expectedLen int
	}{
		{
			name:        "Empty repository",
			expectedLen: 0,
		},
		{
			name:        "Get all categories",
			categories:  repository.Categories{{Name: "Category 1"}, {Name: "Category 2"}},
			wantResult:  repository.Categories{{ID: 1, Name: "Category 1"}, {ID: 2, Name: "Category 2"}},
			expectedLen: 2,
		},
	}

	repo := memory.NewCategoryRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, category := range tc.categories {
				_, err := repo.Create(ctx, category)
				assert.NoError(t, err)
			}

			categories, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, categories)
			assert.Equal(t, tc.expectedLen, len(categories))
		})
	}
}

func TestCategoryRepository_Update(t *testing.T) {
	testCases := []struct {
		name            string
		initialCategory entity.Category
		updatedCategory entity.Category
		wantResult      *entity.Category
		wantErr         error
	}{
		{
			name:            "Update an existing category",
			initialCategory: entity.Category{Name: "Category 1"},
			updatedCategory: entity.Category{ID: 1, Name: "Updated Category"},
			wantResult:      &entity.Category{ID: 1, Name: "Updated Category"},
			wantErr:         nil,
		},
		{
			name:            "Update an existing category",
			initialCategory: entity.Category{Name: "Category 2"},
			updatedCategory: entity.Category{ID: 2, Name: "Updated Category 2"},
			wantResult:      &entity.Category{ID: 2, Name: "Updated Category 2"},
			wantErr:         nil,
		},
		{
			name:            "Update a non-existing category",
			initialCategory: entity.Category{Name: "Category 3"},
			updatedCategory: entity.Category{ID: 10, Name: "Updated Category 2"},
			wantResult:      nil,
			wantErr:         repository.ErrUpdateCategory,
		},
	}

	repo := memory.NewCategoryRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.initialCategory)
			assert.NoError(t, err)

			result, err := repo.Update(ctx, tc.updatedCategory)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantResult, result)
			}
		})
	}
}

func TestCategoryRepository_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		category entity.Category
		id       uint
		wantErr  error
	}{
		{
			name:     "Delete an existing category",
			category: entity.Category{Name: "Category 1"},
			id:       1,
			wantErr:  nil,
		},
		{
			name:     "Delete an existing category",
			category: entity.Category{Name: "Category 2"},
			id:       1,
			wantErr:  nil,
		},
		{
			name:     "Delete a non-existing category",
			category: entity.Category{Name: "Category 3"},
			id:       10,
			wantErr:  repository.ErrDeleteCategory,
		},
	}

	repo := memory.NewCategoryRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.category)
			assert.NoError(t, err)

			err = repo.Delete(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
