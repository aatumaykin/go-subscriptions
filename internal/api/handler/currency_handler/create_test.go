package currency_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandle(t *testing.T) {
	type requestDTO struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	type responseDTO struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name           string
		requestBody    requestDTO
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "Test Create USD",
			requestBody:    requestDTO{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success(responseDTO{Code: "USD", Name: "US Dollar", Symbol: "$"}),
		},
		{
			name:           "Test Already exists",
			requestBody:    requestDTO{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(repository.ErrAlreadyExistsCurrency),
		},
		{
			name:           "Test Create EUR",
			requestBody:    requestDTO{Code: "EUR", Name: "Euro", Symbol: "€"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success(responseDTO{Code: "EUR", Name: "Euro", Symbol: "€"}),
		},
		{
			name:           "Test validation error",
			requestBody:    requestDTO{},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(service.ErrInvalidCurrency),
		},
		{
			name:           "Test validation error",
			requestBody:    requestDTO{Name: "EUR", Symbol: "€"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(service.ErrInvalidCurrency),
		},
		{
			name:           "Test validation error",
			requestBody:    requestDTO{Code: "EUR", Symbol: "€"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(service.ErrInvalidCurrency),
		},
		{
			name:           "Test validation error",
			requestBody:    requestDTO{Code: "EUR", Name: "Euro"},
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(service.ErrInvalidCurrency),
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)

			w := httptest.NewRecorder()
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}

			currency_handler.CreateHandle(ctx, cs)(w, r, nil)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
