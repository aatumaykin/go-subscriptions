package memory_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestCurrencyRepository_Create(t *testing.T) {
	testCases := []struct {
		name       string
		currency   entity.Currency
		wantResult *entity.Currency
		wantErr    error
	}{
		{
			name:       "Create a new currency",
			currency:   entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			wantResult: &entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			wantErr:    nil,
		},
		{
			name:       "Create a duplicate currency",
			currency:   entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			wantResult: nil,
			wantErr:    repository.ErrAlreadyExistsCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := repo.Create(ctx, tc.currency)

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

func TestCurrencyRepository_Get(t *testing.T) {
	testCases := []struct {
		name     string
		currency entity.Currency
		code     string
		wantErr  error
	}{
		{
			name:     "Currency does not exist",
			currency: entity.Currency{Code: "NOT", Name: "Not Exist", Symbol: "?"},
			code:     "RUB",
			wantErr:  repository.ErrNotFoundCurrency,
		},
		{
			name:     "Currency by code",
			currency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			code:     "USD",
			wantErr:  nil,
		},
		{
			name:     "Currency by code",
			currency: entity.Currency{Code: "RUB", Name: "Russian Ruble", Symbol: "₽"},
			code:     "RUB",
			wantErr:  nil,
		},
	}

	repo := memory.NewCurrencyRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createdCurrency, err := repo.Create(ctx, tc.currency)
			assert.NoError(t, err)

			foundCurrency, err := repo.Get(ctx, tc.code)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, createdCurrency, foundCurrency)
			}
		})
	}
}

func TestCurrencyRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name        string
		currencies  repository.Currencies
		wantResult  repository.Currencies
		expectedLen int
	}{
		{
			name:        "Empty repository",
			expectedLen: 0,
		},
		{
			name:        "Get all currencies",
			currencies:  repository.Currencies{{Code: "USD", Name: "US Dollar", Symbol: "$"}, {Code: "RUB", Name: "Russian Ruble", Symbol: "₽"}},
			wantResult:  repository.Currencies{{Code: "USD", Name: "US Dollar", Symbol: "$"}, {Code: "RUB", Name: "Russian Ruble", Symbol: "₽"}},
			expectedLen: 2,
		},
	}

	repo := memory.NewCurrencyRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, currency := range tc.currencies {
				_, err := repo.Create(ctx, currency)
				assert.NoError(t, err)
			}

			currencies, err := repo.GetAll(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, currencies)
			assert.Equal(t, tc.expectedLen, len(currencies))
		})
	}
}

func TestCurrencyRepository_Update(t *testing.T) {
	testCases := []struct {
		name            string
		initialCurrency entity.Currency
		updatedCurrency entity.Currency
		wantResult      *entity.Currency
		wantErr         error
	}{
		{
			name:            "Update an existing currency",
			initialCurrency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			updatedCurrency: entity.Currency{Code: "USD", Name: "Euro", Symbol: "€"},
			wantResult:      &entity.Currency{Code: "USD", Name: "Euro", Symbol: "€"},
			wantErr:         nil,
		},
		{
			name:            "Update a non-existing currency",
			initialCurrency: entity.Currency{Code: "RUB", Name: "Russian Ruble", Symbol: "₽"},
			updatedCurrency: entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			wantResult:      nil,
			wantErr:         repository.ErrNotFoundCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.initialCurrency)
			assert.NoError(t, err)

			result, err := repo.Update(ctx, tc.updatedCurrency)

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

func TestCurrencyRepository_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		currency entity.Currency
		code     string
		wantErr  error
	}{
		{
			name:     "Delete an existing currency",
			currency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			code:     "USD",
			wantErr:  nil,
		},
		{
			name:    "Delete a non-existing currency",
			code:    "EUR",
			wantErr: repository.ErrNotFoundCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Create(ctx, tc.currency)
			assert.NoError(t, err)

			err = repo.Delete(ctx, tc.code)

			if tc.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
