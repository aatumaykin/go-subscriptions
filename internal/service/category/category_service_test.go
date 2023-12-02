package category_test

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/category/entity"
	"git.home/alex/go-subscriptions/internal/domain/category/repository"
	"git.home/alex/go-subscriptions/internal/service/category"
)

func TestCategoryService_Create(t *testing.T) {
	type testCase struct {
		test             string
		name             string
		expectedCategory *entity.Category
		expectedErr      error
	}

	testCases := []testCase{
		{
			test:             "Create a new category",
			name:             "Test Category 1",
			expectedCategory: &entity.Category{ID: 1, Name: "Test Category 1"},
			expectedErr:      nil,
		},
		{
			test:             "Create a new category",
			name:             "Test Category 2",
			expectedCategory: &entity.Category{ID: 2, Name: "Test Category 2"},
			expectedErr:      nil,
		},
	}

	s, err := category.NewCategoryService(category.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			createdCategory, err := s.Create(context.Background(), tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if createdCategory != nil {
				if *createdCategory != *tc.expectedCategory {
					t.Errorf("Expected %v, got %v", tc.expectedCategory, createdCategory)
				}
			}
		})
	}
}

func TestCategoryService_Get(t *testing.T) {
	type testCase struct {
		test             string
		expectedCategory *entity.Category
		expectedErr      error
	}

	testCases := []testCase{
		{
			test:             "Get a non-existing category",
			expectedCategory: &entity.Category{ID: 1, Name: "Test Category"},
			expectedErr:      repository.ErrNotFoundCategory,
		},
		{
			test:             "Get an existing category",
			expectedCategory: &entity.Category{ID: 1, Name: "Test Category 1"},
			expectedErr:      nil,
		},
	}

	s, err := category.NewCategoryService(category.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if tc.expectedErr == nil {
				_, err := s.Create(context.Background(), tc.expectedCategory.Name)
				if err != nil {
					t.Fatal(err)
				}
			}

			found, err := s.Get(context.Background(), tc.expectedCategory.ID)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if found != nil {
				if *found != *tc.expectedCategory {
					t.Errorf("Expected %v, got %v", tc.expectedCategory, found)
				}
			}
		})
	}
}

func TestCategoryService_GetAll(t *testing.T) {
	type testCase struct {
		test               string
		expectedCategories repository.Categories
		expectedErr        error
	}

	testCases := []testCase{
		{
			test: "Empty list",
		},
		{
			test:               "Get all categories",
			expectedCategories: repository.Categories{{ID: 1, Name: "Test Category 1"}, {ID: 2, Name: "Test Category 2"}},
		},
	}

	s, err := category.NewCategoryService(category.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if tc.expectedErr == nil {
				for _, c := range tc.expectedCategories {
					_, err := s.Create(context.Background(), c.Name)
					if err != nil {
						t.Fatal(err)
					}
				}
			}

			categories, err := s.GetAll(context.Background())
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
			if len(categories) != len(tc.expectedCategories) {
				t.Errorf("Expected %d categories, got %d", len(tc.expectedCategories), len(categories))
			}
			if len(categories) > 0 {
				for i, c := range categories {
					if c.ID != tc.expectedCategories[i].ID || c.Name != tc.expectedCategories[i].Name {
						t.Errorf("Expected %v, got %v", tc.expectedCategories[i], c)
					}
				}
			}
		})
	}
}

func TestCategoryService_Update(t *testing.T) {
	type testCase struct {
		test            string
		name            string
		updatedCategory entity.Category
		expectedErr     error
	}

	testCases := []testCase{
		{
			test:            "Update an existing category",
			name:            "Category 1",
			updatedCategory: entity.Category{ID: 1, Name: "Updated Category 1"},
			expectedErr:     nil,
		},
		{
			test:            "Update an existing category 2",
			name:            "Category 2",
			updatedCategory: entity.Category{ID: 2, Name: "Updated Category 2"},
			expectedErr:     nil,
		},
		{
			test:            "Update a non-existing category",
			updatedCategory: entity.Category{ID: 3, Name: "Updated Category"},
			expectedErr:     repository.ErrUpdateCategory,
		},
	}

	s, err := category.NewCategoryService(category.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if tc.name != "" {
				_, err := s.Create(context.Background(), tc.name)
				if err != nil {
					t.Fatal(err)
				}
			}

			updatedCategory, err := s.Update(context.Background(), tc.updatedCategory)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if updatedCategory != nil {
				if *updatedCategory != tc.updatedCategory {
					t.Errorf("Expected %v, got %v", tc.updatedCategory, updatedCategory)
				}
			}
		})
	}
}

func TestCategoryService_Delete(t *testing.T) {
	s, err := category.NewCategoryService(category.WithMemoryCategoryRepository())
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Delete a non-existing category", func(t *testing.T) {
		categories, err := s.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if len(categories) > 0 {
			t.Errorf("Expected 0 categories, got %d", len(categories))
		}

		err = s.Delete(context.Background(), 1)
		if !errors.Is(err, repository.ErrDeleteCategory) {
			t.Errorf("Expected error: %v, got: %v", repository.ErrDeleteCategory, err)
		}
	})

	t.Run("Delete an existing category", func(t *testing.T) {
		_, err := s.Create(context.Background(), "Category 1")
		if err != nil {
			t.Fatal(err)
		}

		_, err = s.Create(context.Background(), "Category 2")
		if err != nil {
			t.Fatal(err)
		}

		categories, err := s.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(categories) != 2 {
			t.Errorf("Expected 2 categories, got %d", len(categories))
		}

		err = s.Delete(context.Background(), 1)
		if err != nil {
			t.Fatal(err)
		}

		categories, err = s.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(categories) != 1 {
			t.Errorf("Expected 1 categories, got %d", len(categories))
		}

		_, err = s.Get(context.Background(), 1)
		if !errors.Is(err, repository.ErrNotFoundCategory) {
			t.Errorf("Expected error: %v, got: %v", repository.ErrNotFoundCategory, err)
		}

		_, err = s.Get(context.Background(), 2)
		if err != nil {
			t.Fatal(err)
		}

		err = s.Delete(context.Background(), 2)
		if err != nil {
			t.Fatal(err)
		}

		categories, err = s.GetAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if len(categories) != 0 {
			t.Errorf("Expected 0 categories, got %d", len(categories))
		}
	})
}
