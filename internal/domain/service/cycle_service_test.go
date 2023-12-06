package service_test

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
)

type MockCycleRepository struct {
	mock.Mock
}

func (m *MockCycleRepository) Create(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	args := m.Called(ctx, cycle)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Cycle), args.Error(1)
}

func (m *MockCycleRepository) Get(ctx context.Context, ID uint) (*entity.Cycle, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Cycle), args.Error(1)
}

func (m *MockCycleRepository) GetAll(ctx context.Context) (repository.Cycles, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(repository.Cycles), args.Error(1)
}

func (m *MockCycleRepository) Update(ctx context.Context, cycle entity.Cycle) (*entity.Cycle, error) {
	args := m.Called(ctx, cycle)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Cycle), args.Error(1)
}

func (m *MockCycleRepository) Delete(ctx context.Context, ID uint) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}

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
			mockRepo := new(MockCycleRepository)
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
			wantErr:    errors.New("some error"),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockCycleRepository)
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
			wantErr:    errors.New("some error"),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockCycleRepository)
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
			wantErr:    errors.New("some error"),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockCycleRepository)
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
			wantErr: errors.New("some error"),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockCycleRepository)
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
