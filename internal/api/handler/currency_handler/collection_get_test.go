package currency_handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/currency_handler"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/tests"
	"git.home/alex/go-subscriptions/tests/mock_repository"
	"github.com/stretchr/testify/assert"
)

func TestCollectionGetHandle(t *testing.T) {
	type responseDTO struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	testCases := []struct {
		name           string
		currencies     repository.Currencies
		mockError      error
		expectedStatus int
		expectedBody   api_response.ResponseDTO
	}{
		{
			name:           "Empty currencies",
			currencies:     repository.Currencies{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Success([]responseDTO{}),
		},
		{
			name: "Success",
			currencies: repository.Currencies{
				{Code: "USD", Symbol: "$", Name: "US Dollar"},
				{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: api_response.Success([]responseDTO{
				{Code: "USD", Symbol: "$", Name: "US Dollar"},
				{Code: "RUB", Symbol: "₽", Name: "Russian Ruble"},
			}),
		},
		{
			name:           "Error",
			currencies:     nil,
			mockError:      tests.ErrTest,
			expectedStatus: http.StatusOK,
			expectedBody:   api_response.Error(tests.ErrTest),
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mock_repository.MockCurrencyRepository)
			mockRepo.On("GetAll", context.Background()).Return(tc.currencies, tc.mockError)

			cs := service.NewCurrencyService(mockRepo)
			w := httptest.NewRecorder()

			currency_handler.CollectionGetHandle(ctx, cs)(w, nil, nil)

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "handler returned wrong content type")

			expectedBody, err := json.Marshal(tc.expectedBody)
			assert.NoError(t, err)

			assert.Equal(t, string(expectedBody), w.Body.String())
		})
	}
}
