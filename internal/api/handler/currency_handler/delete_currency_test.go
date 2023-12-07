package currency_handler_test

import (
	"context"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCurrency(t *testing.T) {
	testCases := []struct {
		name     string
		code     string
		currency entity.Currency
		wantErr  error
	}{
		{
			name:     "success",
			code:     "USD",
			currency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			wantErr:  nil,
		},
		{
			name:     "error",
			code:     "USD",
			currency: entity.Currency{Code: "USD"},
			wantErr:  repository.ErrNotFoundCurrency,
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCurrency(ctx, tc.currency)
			ps := httprouter.Params{{Key: "code", Value: tc.code}}

			response := currency_handler.DeleteCurrency(ctx, cs)(nil, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			assert.Nil(t, response)
		})
	}
}
