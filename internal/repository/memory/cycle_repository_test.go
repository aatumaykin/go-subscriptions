package memory_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/stretchr/testify/assert"

	"git.home/alex/go-subscriptions/internal/domain/entity"
)

func TestCycleRepository_Create(t *testing.T) {
	testCases := []struct {
		name       string
		cycle      entity.Cycle
		wantResult *entity.Cycle
		wantErr    error
	}{
		{
			name:       "Create a new cycle",
			cycle:      entity.Cycle{Name: "Test Cycle"},
			wantResult: &entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantErr:    nil,
		},
		{
			name:       "Create a new cycle",
			cycle:      entity.Cycle{Name: "Test Cycle"},
			wantResult: &entity.Cycle{ID: 2, Name: "Test Cycle"},
			wantErr:    nil,
		},
	}

	repo := memory.NewCycleRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Create(ctx, tc.cycle)

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

func TestCycleRepository_Get(t *testing.T) {
	type testCase struct {
		name    string
		cycle   entity.Cycle
		id      uint
		wantErr error
	}

	testCases := []testCase{
		{
			name:    "Get an existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle"},
			id:      1,
			wantErr: nil,
		},
		{
			name:    "Get an existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle"},
			id:      2,
			wantErr: nil,
		},
		{
			name:    "Get a non-existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle"},
			id:      10,
			wantErr: repository.ErrNotFoundCycle,
		},
	}

	repo := memory.NewCycleRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createdCycle, err := repo.Create(ctx, tc.cycle)
			assert.NoError(t, err)

			foundCycle, err := repo.Get(ctx, tc.id)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, createdCycle, foundCycle)
			}
		})
	}
}

func TestCycleRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		cycles      repository.Cycles
		wantResult  repository.Cycles
		expectedLen int
	}{
		{
			name:        "Empty repository",
			expectedLen: 0,
		},
		{
			name:        "Get all categories",
			cycles:      repository.Cycles{{Name: "Cycle 1"}, {Name: "Cycle 2"}},
			wantResult:  repository.Cycles{{ID: 1, Name: "Cycle 1"}, {ID: 2, Name: "Cycle 2"}},
			expectedLen: 2,
		},
	}

	repo := memory.NewCycleRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, cycle := range tc.cycles {
				_, err := repo.Create(ctx, cycle)
				assert.NoError(t, err)
			}

			cycles, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, cycles)
			assert.Equal(t, tc.expectedLen, len(cycles))
		})
	}
}

func TestCycleRepository_Update(t *testing.T) {
	testCases := []struct {
		name         string
		initialCycle entity.Cycle
		updatedCycle entity.Cycle
		wantResult   *entity.Cycle
		wantErr      error
	}{
		{
			name:         "Update an existing cycle",
			initialCycle: entity.Cycle{Name: "Cycle 1"},
			updatedCycle: entity.Cycle{ID: 1, Name: "Updated Cycle"},
			wantResult:   &entity.Cycle{ID: 1, Name: "Updated Cycle"},
			wantErr:      nil,
		},
		{
			name:         "Update an existing cycle",
			initialCycle: entity.Cycle{Name: "Cycle 2"},
			updatedCycle: entity.Cycle{ID: 2, Name: "Updated Cycle"},
			wantResult:   &entity.Cycle{ID: 2, Name: "Updated Cycle"},
			wantErr:      nil,
		},
		{
			name:         "Update a non-existing cycle",
			initialCycle: entity.Cycle{Name: "Cycle 2"},
			updatedCycle: entity.Cycle{ID: 10, Name: "Updated Cycle"},
			wantErr:      repository.ErrUpdateCycle,
		},
	}

	repo := memory.NewCycleRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.initialCycle)
			assert.NoError(t, err)

			result, err := repo.Update(ctx, tc.updatedCycle)

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

func TestCycleRepository_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		cycle   entity.Cycle
		id      uint
		wantErr error
	}{
		{
			name:    "Delete an existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle"},
			id:      1,
			wantErr: nil,
		},
		{
			name:    "Delete an existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle 2"},
			id:      1,
			wantErr: nil,
		},
		{
			name:    "Delete a non-existing cycle",
			cycle:   entity.Cycle{Name: "Test Cycle 3"},
			id:      10,
			wantErr: repository.ErrDeleteCycle,
		},
	}

	repo := memory.NewCycleRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.cycle)
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
