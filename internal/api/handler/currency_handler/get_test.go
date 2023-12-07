package currency_handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestGetHandle(t *testing.T) {
	type responseDTO struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name           string
		code           string
		currency       entity.Currency
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "success",
			code:           "USD",
			currency:       entity.Currency{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success(responseDTO{Code: "USD", Name: "US Dollar", Symbol: "$"}),
		},
		{
			name:           "Test not found",
			code:           "RUB",
			currency:       entity.Currency{Code: "EUR", Name: "Euro", Symbol: "€"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(repository.ErrNotFoundCurrency),
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, _ = cs.CreateCurrency(ctx, tc.currency)

			w := httptest.NewRecorder()
			ps := httprouter.Params{{Key: "code", Value: tc.code}}

			currency_handler.GetHandle(ctx, cs)(w, nil, ps)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
