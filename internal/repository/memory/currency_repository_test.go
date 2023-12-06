package memory_test

import (
	"context"
	"errors"
	"testing"

	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/repository/memory"
)

func TestCurrencyRepository_Create(t *testing.T) {
	type testCase struct {
		test        string
		currency    entity.Currency
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Create a new currency",
			currency:    entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedErr: nil,
		},
		{
			test:        "Create a duplicate currency",
			currency:    entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedErr: repository.ErrCreateCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			createdCurrency, err := repo.Create(context.Background(), tc.currency)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if createdCurrency != nil {
				found, err := repo.Get(context.Background(), createdCurrency.Code)
				if err != nil {
					t.Fatal(err)
				}
				if createdCurrency.Name != found.Name {
					t.Errorf("Expected %v, got %v", createdCurrency.Name, found.Name)
				}
				if createdCurrency.Symbol != found.Symbol {
					t.Errorf("Expected %v, got %v", createdCurrency.Symbol, found.Symbol)
				}
				if createdCurrency.Code != found.Code {
					t.Errorf("Expected %v, got %v", createdCurrency.Code, found.Code)
				}
			}
		})
	}
}

func TestCurrencyRepository_Get(t *testing.T) {
	type testCase struct {
		test         string
		currency     entity.Currency
		expectedCode string
		expectedErr  error
	}

	testCases := []testCase{
		{
			test:         "Currency does not exist",
			expectedCode: "RUB",
			expectedErr:  repository.ErrNotFoundCurrency,
		},
		{
			test:         "Currency by code",
			currency:     entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedCode: "USD",
			expectedErr:  nil,
		},
		{
			test:         "Currency by code",
			currency:     entity.Currency{Code: "RUB", Name: "Russian Ruble", Symbol: "₽"},
			expectedCode: "RUB",
			expectedErr:  nil,
		},
	}

	repo := memory.NewCurrencyRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if tc.currency.Code != "" {
				_, err := repo.Create(context.Background(), tc.currency)
				if err != nil {
					t.Fatal(err)
				}
			}

			_, err := repo.Get(context.Background(), tc.expectedCode)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCurrencyRepository_GetAll(t *testing.T) {
	type testCase struct {
		test        string
		currencies  repository.Currencies
		expectedLen int
	}

	testCases := []testCase{
		{
			test:        "Empty repository",
			expectedLen: 0,
		},
		{
			test:        "Get all currencies",
			currencies:  []entity.Currency{{Code: "USD", Name: "US Dollar", Symbol: "$"}, {Code: "RUB", Name: "Russian Ruble", Symbol: "₽"}},
			expectedLen: 2,
		},
	}

	repo := memory.NewCurrencyRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			if len(tc.currencies) > 0 {
				for _, currency := range tc.currencies {
					_, err := repo.Create(context.Background(), currency)
					if err != nil {
						t.Fatal(err)
					}
				}
			}

			currencies, _ := repo.GetAll(context.Background())
			if len(currencies) != len(tc.currencies) {
				t.Errorf("Expected %d currencies, got %v", len(tc.currencies), len(currencies))
			}
		})
	}
}

func TestCurrencyRepository_Update(t *testing.T) {
	type testCase struct {
		test            string
		initialCurrency entity.Currency
		updatedCurrency entity.Currency
		expectedErr     error
	}

	testCases := []testCase{
		{
			test:            "Update an existing currency",
			initialCurrency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			updatedCurrency: entity.Currency{Code: "USD", Name: "Euro", Symbol: "€"},
			expectedErr:     nil,
		},
		{
			test:            "Update a non-existing currency",
			initialCurrency: entity.Currency{Code: "RUB", Name: "Russian Ruble", Symbol: "₽"},
			updatedCurrency: entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			expectedErr:     repository.ErrUpdateCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Create the initial currency in the repository
			_, err := repo.Create(context.Background(), tc.initialCurrency)
			if err != nil {
				t.Fatal(err)
			}

			// Call the Update method
			updatedCurrency, err := repo.Update(context.Background(), tc.updatedCurrency)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			if updatedCurrency != nil {
				// Retrieve the updated currency from the repository
				retrievedCurrency, err := repo.Get(context.Background(), updatedCurrency.Code)
				if err != nil {
					t.Fatal(err)
				}

				// Check if the updated currency's fields match the expected values
				if retrievedCurrency.Code != tc.updatedCurrency.Code {
					t.Errorf("Expected code: %s, got: %s", tc.updatedCurrency.Code, retrievedCurrency.Code)
				}
				if retrievedCurrency.Name != tc.updatedCurrency.Name {
					t.Errorf("Expected code: %s, got: %s", tc.updatedCurrency.Name, retrievedCurrency.Name)
				}
				if retrievedCurrency.Symbol != tc.updatedCurrency.Symbol {
					t.Errorf("Expected code: %s, got: %s", tc.updatedCurrency.Symbol, retrievedCurrency.Symbol)
				}
			}
		})
	}
}

func TestCurrencyRepository_Delete(t *testing.T) {
	type testCase struct {
		test            string
		initialCurrency entity.Currency
		initialCode     string
		expectedErr     error
	}

	testCases := []testCase{
		{
			test:            "Delete an existing currency",
			initialCurrency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			initialCode:     "USD",
			expectedErr:     nil,
		},
		{
			test:        "Delete a non-existing currency",
			initialCode: "EUR",
			expectedErr: repository.ErrDeleteCurrency,
		},
	}

	repo := memory.NewCurrencyRepository()

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			// Create the initial currency in the repository
			if tc.initialCurrency.Code != "" {
				_, err := repo.Create(context.Background(), tc.initialCurrency)
				if err != nil {
					t.Fatal(err)
				}
			}

			// Call the Delete method
			err := repo.Delete(context.Background(), tc.initialCode)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}

			// Check if the currency was deleted
			_, err = repo.Get(context.Background(), tc.initialCode)
			if !errors.Is(err, repository.ErrNotFoundCurrency) {
				t.Errorf("Expected error: %v, got: %v", repository.ErrNotFoundCurrency, err)
			}
		})
	}
}
