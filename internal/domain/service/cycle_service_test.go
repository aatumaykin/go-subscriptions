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

func TestCycleService_CreateCycle(t *testing.T) {
	testCases := []struct {
		name       string
		cycle      entity.Cycle
		wantResult *entity.Cycle
		wantErr    error
	}{
		{
			name:       "Test empty cycle name",
			cycle:      entity.Cycle{Name: ""},
			wantResult: nil,
			wantErr:    service.ErrInvalidCycle,
		},
		{
			name:       "Test valid cycle",
			cycle:      entity.Cycle{Name: "Test Cycle"},
			wantResult: &entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			cycle:      entity.Cycle{Name: "Test Cycle"},
			wantResult: &entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantErr:    nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("Create", ctx, tc.cycle).Return(tc.wantResult, tc.wantErr)

			cycleService := service.NewCycleService(mockRepo)
			result, err := cycleService.CreateCycle(ctx, tc.cycle)

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

func TestCycleService_GetCycle(t *testing.T) {
	testCases := []struct {
		name       string
		id         uint
		wantResult *entity.Cycle
		wantErr    error
	}{
		{
			name:       "ID is zero",
			id:         0,
			wantResult: nil,
			wantErr:    repository.ErrNotFoundCycle,
		},
		{
			name:       "Test valid cycle",
			id:         1,
			wantResult: &entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			id:         1,
			wantResult: nil,
			wantErr:    tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("Get", ctx, tc.id).Return(tc.wantResult, tc.wantErr)

			cycleService := service.NewCycleService(mockRepo)
			result, err := cycleService.GetCycle(ctx, tc.id)

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

func TestCycleService_GetAllCycles(t *testing.T) {
	testCases := []struct {
		name       string
		wantResult repository.Cycles
		wantErr    error
	}{
		{
			name:       "Test empty cycles",
			wantResult: repository.Cycles{},
			wantErr:    nil,
		},
		{
			name:       "Test valid cycles",
			wantResult: repository.Cycles{{ID: 1, Name: "Test Cycle"}},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			wantResult: nil,
			wantErr:    tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("GetAll", ctx).Return(tc.wantResult, tc.wantErr)

			cycleService := service.NewCycleService(mockRepo)
			result, err := cycleService.GetAllCycles(ctx)

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

func TestCycleService_UpdateCycle(t *testing.T) {
	testCases := []struct {
		name       string
		cycle      entity.Cycle
		wantResult *entity.Cycle
		wantErr    error
	}{
		{
			name:       "ID is zero",
			cycle:      entity.Cycle{ID: 0, Name: "Test Cycle"},
			wantResult: nil,
			wantErr:    service.ErrInvalidCycle,
		},
		{
			name:       "Name is an empty string",
			cycle:      entity.Cycle{ID: 1, Name: ""},
			wantResult: nil,
			wantErr:    service.ErrInvalidCycle,
		},
		{
			name:       "Test valid cycle",
			cycle:      entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantResult: &entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			cycle:      entity.Cycle{ID: 1, Name: "Test Cycle"},
			wantResult: nil,
			wantErr:    tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("Update", ctx, tc.cycle).Return(tc.wantResult, tc.wantErr)

			cycleService := service.NewCycleService(mockRepo)
			result, err := cycleService.UpdateCycle(ctx, tc.cycle)

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

func TestCycleService_DeleteCycle(t *testing.T) {
	testCases := []struct {
		name    string
		id      uint
		wantErr error
	}{
		{
			name:    "ID is zero",
			id:      0,
			wantErr: repository.ErrNotFoundCycle,
		},
		{
			name:    "Test valid cycle",
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
			mockRepo := new(mock_repository.MockCycleRepository)
			mockRepo.On("Delete", ctx, tc.id).Return(tc.wantErr)

			cycleService := service.NewCycleService(mockRepo)
			err := cycleService.DeleteCycle(ctx, tc.id)

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
