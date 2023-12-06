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

func TestCurrencyService_CreateCurrency(t *testing.T) {
	testCases := []struct {
		name       string
		currency   entity.Currency
		wantResult *entity.Currency
		wantErr    error
	}{
		{
			name:       "Test valid currency",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantResult: &entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantErr:    nil,
		},
		{
			name:       "Test empty currency code",
			currency:   entity.Currency{Code: "", Symbol: "$", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test empty currency symbol",
			currency:   entity.Currency{Code: "USD", Symbol: "", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test empty currency name",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: ""},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test error",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    repository.ErrCreateCurrency,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("Create", ctx, tc.currency).Return(tc.wantResult, tc.wantErr)

			currencyService := service.NewCurrencyService(mockRepo)
			result, err := currencyService.CreateCurrency(ctx, tc.currency)

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

func TestCurrencyService_GetCurrency(t *testing.T) {
	testCases := []struct {
		name       string
		code       string
		wantResult *entity.Currency
		wantErr    error
	}{
		{
			name:       "Test valid currency",
			code:       "USD",
			wantResult: &entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantErr:    nil,
		},
		{
			name:       "Test empty currency code",
			code:       "",
			wantResult: nil,
			wantErr:    repository.ErrNotFoundCurrency,
		},
		{
			name:       "Test error",
			code:       "USD",
			wantResult: nil,
			wantErr:    repository.ErrNotFoundCurrency,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("Get", ctx, tc.code).Return(tc.wantResult, tc.wantErr)

			currencyService := service.NewCurrencyService(mockRepo)
			result, err := currencyService.GetCurrency(ctx, tc.code)

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

func TestCurrencyService_GetAllCurrencies(t *testing.T) {
	testCases := []struct {
		name       string
		wantResult repository.Currencies
		wantErr    error
	}{
		{
			name:       "Test empty currencies",
			wantResult: repository.Currencies{},
			wantErr:    nil,
		},
		{
			name: "Test valid currencies",
			wantResult: repository.Currencies{
				{Code: "USD", Symbol: "$", Name: "US Dollar"},
				{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"},
			},
			wantErr: nil,
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
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("GetAll", ctx).Return(tc.wantResult, tc.wantErr)

			currencyService := service.NewCurrencyService(mockRepo)
			result, err := currencyService.GetAllCurrencies(ctx)

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

func TestCurrencyService_UpdateCurrency(t *testing.T) {
	testCases := []struct {
		name       string
		currency   entity.Currency
		wantResult *entity.Currency
		wantErr    error
	}{
		{
			name:       "Test valid currency",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantResult: &entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantErr:    nil,
		},
		{
			name:       "Test empty currency code",
			currency:   entity.Currency{Code: "", Symbol: "$", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test empty currency symbol",
			currency:   entity.Currency{Code: "USD", Symbol: "", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test empty currency name",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: ""},
			wantResult: nil,
			wantErr:    service.ErrInvalidCurrency,
		},
		{
			name:       "Test error",
			currency:   entity.Currency{Code: "USD", Symbol: "$", Name: "US Dollar"},
			wantResult: nil,
			wantErr:    repository.ErrUpdateCurrency,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("Update", ctx, tc.currency).Return(tc.wantResult, tc.wantErr)

			currencyService := service.NewCurrencyService(mockRepo)
			result, err := currencyService.UpdateCurrency(ctx, tc.currency)

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

func TestCurrencyService_DeleteCurrency(t *testing.T) {
	testCases := []struct {
		name    string
		code    string
		wantErr error
	}{
		{
			name:    "Test valid currency",
			code:    "USD",
			wantErr: nil,
		},
		{
			name:    "Test empty currency code",
			code:    "",
			wantErr: repository.ErrNotFoundCurrency,
		},
		{
			name:    "Test error",
			code:    "USD",
			wantErr: repository.ErrDeleteCurrency,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("Delete", ctx, tc.code).Return(tc.wantErr)

			currencyService := service.NewCurrencyService(mockRepo)
			err := currencyService.DeleteCurrency(ctx, tc.code)

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
