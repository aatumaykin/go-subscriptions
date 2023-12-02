package memory

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/cycle/repository"

	"git.home/alex/go-subscriptions/internal/domain/cycle/entity"
)

func TestCycleRepository_Create(t *testing.T) {
	type testCase struct {
		test        string
		cycle       entity.Cycle
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Create a new cycle",
			cycle:       entity.Cycle{Name: "Test Cycle"},
			expectedErr: nil,
		},
		{
			test:        "Create a new cycle",
			cycle:       entity.Cycle{ID: 1, Name: "Test Cycle"},
			expectedErr: nil,
		},
	}

	repo := NewCycleRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			createdCycle, err := repo.Create(context.Background(), tc.cycle)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if createdCycle != nil {
				found, err := repo.GetByID(context.Background(), createdCycle.ID)
				if err != nil {
					t.Fatal(err)
				}
				if createdCycle.Name != found.Name {
					t.Errorf("Expected %v, got %v", createdCycle.Name, found.Name)
				}
			}
		})
	}
}

func TestCycleRepository_GetByID(t *testing.T) {
	type testCase struct {
		test        string
		cycle       entity.Cycle
		expectedID  uint
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Get an existing cycle",
			cycle:       entity.Cycle{Name: "Test Cycle"},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			test:        "Get a non-existing cycle",
			cycle:       entity.Cycle{Name: "Test Cycle"},
			expectedID:  3,
			expectedErr: repository.ErrNotFoundCycle,
		},
	}

	repo := NewCycleRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.cycle)
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

func TestCycleRepository_GetAll(t *testing.T) {
	testCases := []struct {
		test        string
		categories  []entity.Cycle
		expectedLen int
	}{
		{
			test:        "Empty repository",
			expectedLen: 0,
		},
		{
			test:        "Get all categories",
			categories:  []entity.Cycle{{Name: "Cycle 1"}, {Name: "Cycle 2"}},
			expectedLen: 2,
		},
	}

	repo := NewCycleRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			for _, cycle := range tc.categories {
				_, err := repo.Create(context.Background(), cycle)
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

func TestCycleRepository_Update(t *testing.T) {
	testCases := []struct {
		test         string
		initialCycle entity.Cycle
		updatedCycle entity.Cycle
		expectedErr  error
	}{
		{
			test:         "Update an existing cycle",
			initialCycle: entity.Cycle{Name: "Cycle 1"},
			updatedCycle: entity.Cycle{ID: 1, Name: "Updated Cycle"},
			expectedErr:  nil,
		},
		{
			test:         "Update a non-existing cycle",
			initialCycle: entity.Cycle{Name: "Cycle 2"},
			updatedCycle: entity.Cycle{ID: 3, Name: "Updated Cycle"},
			expectedErr:  repository.ErrUpdateCycle,
		},
	}

	repo := NewCycleRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Create(context.Background(), tc.initialCycle)
			if err != nil {
				t.Fatal(err)
			}

			updatedCycle, err := repo.Update(context.Background(), tc.updatedCycle)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if err == nil {
				found, err := repo.GetByID(context.Background(), updatedCycle.ID)
				if err != nil {
					t.Fatal(err)
				}

				if found.Name != updatedCycle.Name {
					t.Errorf("Expected %v, got %v", updatedCycle.Name, found.Name)
				}
			}
		})
	}
}

func TestCycleRepository_Delete(t *testing.T) {
	repo := NewCycleRepository()

	ctx := context.Background()

	t.Run("Non-existing cycle", func(t *testing.T) {
		err := repo.Delete(ctx, 3)
		if err == nil {
			t.Error("Expected error, got nil")
		}

		if !errors.Is(err, repository.ErrDeleteCycle) {
			t.Errorf("Expected error: %v, got: %v", repository.ErrDeleteCycle, err)
		}
	})

	t.Run("Existing cycle", func(t *testing.T) {
		_, err := repo.Create(context.Background(), entity.Cycle{Name: "Cycle 1"})
		if err != nil {
			t.Fatal(err)
		}

		err = repo.Delete(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if _, err := repo.GetByID(context.Background(), 1); err == nil {
			t.Errorf("Expected cycle to be deleted")
		}
	})
}
