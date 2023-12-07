package currency_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCurrency(t *testing.T) {
	type req struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	type resp struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name            string
		initialCurrency entity.Currency
		requestBody     req
		code            string
		expected        resp
		wantErr         error
	}{
		{
			name:            "success",
			initialCurrency: entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			requestBody:     req{Name: "Euro", Symbol: "€"},
			code:            "USD",
			expected:        resp{Code: "USD", Name: "Euro", Symbol: "€"},
		},
		{
			name:            "validation error",
			initialCurrency: entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			requestBody:     req{Name: "", Symbol: "$"},
			code:            "EUR",
			wantErr:         service.ErrInvalidCurrency,
		},
		{
			name:            "validation error",
			initialCurrency: entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			requestBody:     req{Name: "USD", Symbol: ""},
			code:            "EUR",
			wantErr:         service.ErrInvalidCurrency,
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCurrency(ctx, tc.initialCurrency)

			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}
			ps := httprouter.Params{{Key: "code", Value: tc.code}}

			response := currency_handler.UpdateCurrency(ctx, cs)(r, ps)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
