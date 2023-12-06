package mock_repository

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"github.com/stretchr/testify/mock"
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

func (m *MockCycleRepository) Get(ctx context.Context, id uint) (*entity.Cycle, error) {
	args := m.Called(ctx, id)
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

func (m *MockCycleRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
