package memory

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/category/repository"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
)

func TestCategoryRepository_Create(t *testing.T) {
	type testCase struct {
		test        string
		category    entity.Category
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Create a new category",
			category:    entity.Category{Name: "Test Category"},
			expectedErr: nil,
		},
		{
			test:        "Create a new category",
			category:    entity.Category{ID: 1, Name: "Test Category"},
			expectedErr: nil,
		},
	}

	repo := NewCategoryRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			createdCategory, err := repo.Create(context.Background(), tc.category)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if createdCategory != nil {
				found, err := repo.GetByID(context.Background(), createdCategory.ID)
				if err != nil {
					t.Fatal(err)
				}
				if createdCategory.Name != found.Name {
					t.Errorf("Expected %v, got %v", createdCategory.Name, found.Name)
				}
			}
		})
	}
}

func TestCategoryRepository_GetByID(t *testing.T) {
	type testCase struct {
		test        string
		category    entity.Category
		expectedID  uint
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Get an existing category",
			category:    entity.Category{Name: "Test Category"},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			test:        "Get a non-existing category",
			category:    entity.Category{Name: "Test Category"},
			expectedID:  3,
			expectedErr: repository.ErrNotFoundCategory,
		},
	}

	repo := NewCategoryRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.category)
			if err != nil {
				t.Fatal(err)
			}

			_, err = repo.GetByID(context.Background(), tc.expectedID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCategoryRepository_GetAll(t *testing.T) {
	testCases := []struct {
		test        string
		categories  []entity.Category
		expectedLen int
	}{
		{
			test:        "Empty repository",
			expectedLen: 0,
		},
		{
			test:        "Get all categories",
			categories:  []entity.Category{{Name: "Category 1"}, {Name: "Category 2"}},
			expectedLen: 2,
		},
	}

	repo := NewCategoryRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			for _, category := range tc.categories {
				_, err := repo.Create(context.Background(), category)
				if err != nil {
					t.Fatal(err)
				}
			}

			categories, err := repo.GetAll(context.Background())
			if err != nil {
				t.Fatal(err)
			}

			if len(categories) != tc.expectedLen {
				t.Errorf("Expected %d categories, got %d", tc.expectedLen, len(categories))
			}
		})
	}
}

func TestCategoryRepository_Update(t *testing.T) {
	testCases := []struct {
		test            string
		initialCategory entity.Category
		updatedCategory entity.Category
		expectedErr     error
	}{
		{
			test:            "Update an existing category",
			initialCategory: entity.Category{Name: "Category 1"},
			updatedCategory: entity.Category{ID: 1, Name: "Updated Category"},
			expectedErr:     nil,
		},
		{
			test:            "Update a non-existing category",
			initialCategory: entity.Category{Name: "Category 2"},
			updatedCategory: entity.Category{ID: 3, Name: "Updated Category"},
			expectedErr:     repository.ErrUpdateCategory,
		},
	}

	repo := NewCategoryRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.initialCategory)
			if err != nil {
				t.Fatal(err)
			}

			updatedCategory, err := repo.Update(context.Background(), tc.updatedCategory)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if err == nil {
				found, err := repo.GetByID(context.Background(), updatedCategory.ID)
				if err != nil {
					t.Fatal(err)
				}

				if found.Name != updatedCategory.Name {
					t.Errorf("Expected %v, got %v", updatedCategory.Name, found.Name)
				}
			}
		})
	}
}

func TestCategoryRepository_Delete(t *testing.T) {
	repo := NewCategoryRepository()

	ctx := context.Background()

	t.Run("Non-existing category", func(t *testing.T) {
		err := repo.Delete(ctx, 3)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, repository.ErrDeleteCategory) {
			t.Errorf("Expected error: %v, got: %v", repository.ErrDeleteCategory, err)
		}
	})

	t.Run("Existing category", func(t *testing.T) {
		_, err := repo.Create(context.Background(), entity.Category{Name: "Category 1"})
		if err != nil {
			t.Fatal(err)
		}

		err = repo.Delete(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if _, ok := repo.categories[1]; ok {
			t.Errorf("Expected category to be deleted")
		}
	})
}
