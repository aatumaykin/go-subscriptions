package currency_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrency(t *testing.T) {
	type resp struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name     string
		code     string
		currency entity.Currency
		expected resp
		wantErr  error
	}{
		{
			name:     "success",
			code:     "USD",
			currency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expected: resp{Code: "USD", Name: "US Dollar", Symbol: "$"},
		},
		{
			name:     "Test not found",
			code:     "RUB",
			currency: entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			wantErr:  repository.ErrNotFoundCurrency,
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCurrency(ctx, tc.currency)

			ps := httprouter.Params{{Key: "code", Value: tc.code}}

			response := currency_handler.GetCurrency(ctx, cs)(nil, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
