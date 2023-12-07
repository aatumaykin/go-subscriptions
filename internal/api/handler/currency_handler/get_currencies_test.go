package currency_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencies(t *testing.T) {
	type resp struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name       string
		currencies repository.Currencies
		mockError  error
		expected   []resp
		wantErr    error
	}{
		{
			name:       "Empty currencies",
			currencies: repository.Currencies{},
			mockError:  nil,
			expected:   []resp{},
		},
		{
			name: "Success",
			currencies: repository.Currencies{
				{Code: "USD", Symbol: "$", Name: "US Dollar"},
				{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"},
			},
			mockError: nil,
			expected: []resp{
				{Code: "USD", Symbol: "$", Name: "US Dollar"},
				{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"},
			},
		},
		{
			name:       "Error",
			currencies: nil,
			mockError:  tests.ErrTest,
			wantErr:    tests.ErrTest,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("GetAll", ctx).Return(tc.currencies, tc.mockError)

			cs := service.NewCurrencyService(mockRepo)

			response := currency_handler.GetCurrencies(ctx, cs)(nil, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
