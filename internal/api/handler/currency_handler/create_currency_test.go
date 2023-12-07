package currency_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestCreateCurrency(t *testing.T) {
	type req struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	type resp struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name        string
		requestBody req
		expected    resp
		wantErr     error
	}{
		{
			name:        "Test Create USD",
			requestBody: req{Code: "USD", Name: "US Dollar", Symbol: "$"},
			expected:    resp{Code: "USD", Name: "US Dollar", Symbol: "$"},
		},
		{
			name:        "Test Already exists",
			requestBody: req{Code: "USD", Name: "US Dollar", Symbol: "$"},
			wantErr:     repository.ErrAlreadyExistsCurrency,
		},
		{
			name:        "Test Create EUR",
			requestBody: req{Code: "EUR", Name: "Euro", Symbol: "€"},
			expected:    resp{Code: "EUR", Name: "Euro", Symbol: "€"},
		},
		{
			name:        "Test validation error",
			requestBody: req{},
			wantErr:     service.ErrInvalidCurrency,
		},
		{
			name:        "Test validation error",
			requestBody: req{Name: "EUR", Symbol: "€"},
			wantErr:     service.ErrInvalidCurrency,
		},
		{
			name:        "Test validation error",
			requestBody: req{Code: "EUR", Symbol: "€"},
			wantErr:     service.ErrInvalidCurrency,
		},
		{
			name:        "Test validation error",
			requestBody: req{Code: "EUR", Name: "Euro"},
			wantErr:     service.ErrInvalidCurrency,
		},
	}

	cs := service.NewCurrencyService(memory.NewCurrencyRepository())
	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}

			response := currency_handler.CreateCurrency(ctx, cs)(r, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
