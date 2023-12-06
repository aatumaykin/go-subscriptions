package mock_repository

import (
	"context"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"github.com/stretchr/testify/mock"
)

type MockCurrencyRepository struct {
	mock.Mock
}

func (m *MockCurrencyRepository) Create(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	args := m.Called(ctx, currency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Currency), args.Error(1)
}

func (m *MockCurrencyRepository) Get(ctx context.Context, code string) (*entity.Currency, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Currency), args.Error(1)
}

func (m *MockCurrencyRepository) GetAll(ctx context.Context) (repository.Currencies, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(repository.Currencies), args.Error(1)
}

func (m *MockCurrencyRepository) Update(ctx context.Context, currency entity.Currency) (*entity.Currency, error) {
	args := m.Called(ctx, currency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Currency), args.Error(1)
}

func (m *MockCurrencyRepository) Delete(ctx context.Context, code string) error {
	args := m.Called(ctx, code)
	return args.Error(0)
}
