package service_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSubscriptionRepository struct {
	mock.Mock
}

func (m *MockSubscriptionRepository) Create(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	args := m.Called(ctx, subscription)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) Get(ctx context.Context, ID uint) (*entity.Subscription, error) {
	args := m.Called(ctx, ID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) GetAll(ctx context.Context) (repository.Subscriptions, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(repository.Subscriptions), args.Error(1)
}

func (m *MockSubscriptionRepository) Update(ctx context.Context, subscription entity.Subscription) (*entity.Subscription, error) {
	args := m.Called(ctx, subscription)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Subscription), args.Error(1)
}

func (m *MockSubscriptionRepository) Delete(ctx context.Context, ID uint) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}

func TestSubscriptionService_CreateSubscription(t *testing.T) {
	testCases := []struct {
		name         string
		subscription entity.Subscription
		wantResult   *entity.Subscription
		wantErr      error
	}{
		{
			name: "Test empty subscription name",
			subscription: entity.Subscription{
				Name:     "",
				Price:    0,
				Category: entity.Category{ID: 0},
				Currency: entity.Currency{Code: ""},
				Cycle:    entity.Cycle{ID: 0},
			},
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
		{
			name: "Test empty subscription price",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    0,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
		{
			name: "Test empty subscription category",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    100,
				Category: entity.Category{},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: &entity.Subscription{
				ID:       1,
				Name:     "Test",
				Price:    100,
				Category: entity.Category{},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantErr: nil,
		},
		{
			name: "Test empty subscription currency",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: ""},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
		{
			name: "Test empty subscription cycle",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{},
			},
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
		{
			name: "Valid subscription",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: &entity.Subscription{
				ID:       1,
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantErr: nil,
		},
		{
			name: "Test error",
			subscription: entity.Subscription{
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockSubscriptionRepository)
			mockRepo.On("Create", ctx, tc.subscription).Return(tc.wantResult, tc.wantErr)

			subscriptionService := service.NewSubscriptionService(mockRepo)
			result, err := subscriptionService.CreateSubscription(ctx, tc.subscription)

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

func TestSubscriptionService_GetSubscription(t *testing.T) {
	testCases := []struct {
		name       string
		id         uint
		wantResult *entity.Subscription
		wantErr    error
	}{
		{
			name:       "ID is zero",
			id:         0,
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
		{
			name:       "Test valid subscription",
			id:         1,
			wantResult: &entity.Subscription{ID: 1, Name: "Test"},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			id:         1,
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockSubscriptionRepository)
			mockRepo.On("Get", ctx, tc.id).Return(tc.wantResult, tc.wantErr)

			subscriptionService := service.NewSubscriptionService(mockRepo)
			result, err := subscriptionService.GetSubscription(ctx, tc.id)

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

func TestSubscriptionService_GetAllSubscriptions(t *testing.T) {
	testCases := []struct {
		name       string
		wantResult repository.Subscriptions
		wantErr    error
	}{
		{
			name:       "Test empty subscriptions",
			wantResult: repository.Subscriptions{},
			wantErr:    nil,
		},
		{
			name:       "Test valid subscriptions",
			wantResult: repository.Subscriptions{{ID: 1, Name: "Test"}},
			wantErr:    nil,
		},
		{
			name:       "Test error",
			wantResult: nil,
			wantErr:    service.ErrInvalidSubscription,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockSubscriptionRepository)
			mockRepo.On("GetAll", ctx).Return(tc.wantResult, tc.wantErr)

			subscriptionService := service.NewSubscriptionService(mockRepo)
			result, err := subscriptionService.GetAllSubscriptions(ctx)

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

func TestSubscriptionService_UpdateSubscription(t *testing.T) {
	testCases := []struct {
		name         string
		subscription entity.Subscription
		wantResult   *entity.Subscription
		wantErr      error
	}{
		{
			name:         "ID is zero",
			subscription: entity.Subscription{ID: 0, Name: "Test"},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
		{
			name:         "Name is an empty string",
			subscription: entity.Subscription{ID: 1, Name: ""},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
		{
			name:         "Price is zero",
			subscription: entity.Subscription{ID: 1, Name: "Test"},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
		{
			name:         "Currency is an empty",
			subscription: entity.Subscription{ID: 1, Name: "Test"},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
		{
			name:         "Cycle is an empty",
			subscription: entity.Subscription{ID: 1, Name: "Test"},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
		{
			name: "Test valid subscription",
			subscription: entity.Subscription{
				ID:       1,
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantResult: &entity.Subscription{
				ID:       1,
				Name:     "Test",
				Price:    100,
				Category: entity.Category{ID: 1},
				Currency: entity.Currency{Code: "USD"},
				Cycle:    entity.Cycle{ID: 1},
			},
			wantErr: nil,
		},
		{
			name:         "Test error",
			subscription: entity.Subscription{ID: 1, Name: "Test"},
			wantResult:   nil,
			wantErr:      service.ErrInvalidSubscription,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockSubscriptionRepository)
			mockRepo.On("Update", ctx, tc.subscription).Return(tc.wantResult, tc.wantErr)

			subscriptionService := service.NewSubscriptionService(mockRepo)
			result, err := subscriptionService.UpdateSubscription(ctx, tc.subscription)

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

func TestSubscriptionService_DeleteSubscription(t *testing.T) {
	testCases := []struct {
		name    string
		id      uint
		wantErr error
	}{
		{
			name:    "ID is zero",
			id:      0,
			wantErr: service.ErrInvalidSubscription,
		},
		{
			name:    "Test valid subscription",
			id:      1,
			wantErr: nil,
		},
		{
			name:    "Test error",
			id:      1,
			wantErr: service.ErrInvalidSubscription,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockSubscriptionRepository)
			mockRepo.On("Delete", ctx, tc.id).Return(tc.wantErr)

			subscriptionService := service.NewSubscriptionService(mockRepo)
			err := subscriptionService.DeleteSubscription(ctx, tc.id)

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
