package subscription_handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"git.home/alex/go-subscriptions/internal/api/handler/subscription_handler"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/repository"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"git.home/alex/go-subscriptions/internal/repository/memory"
	"git.home/alex/go-subscriptions/tests/tests_assert"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubscription(t *testing.T) {
	type req struct {
		Name            string  `json:"name"`
		Note            string  `json:"note,omitempty"`
		Logo            string  `json:"logo,omitempty"`
		Price           float64 `json:"price"`
		CategoryID      uint    `json:"category_id"`
		CycleID         uint    `json:"cycle_id"`
		CurrencyCode    string  `json:"currency"`
		NextPaymentDate string  `json:"next_payment_date"`
	}

	type resp struct {
		ID              uint    `json:"id"`
		Name            string  `json:"name"`
		Note            string  `json:"note"`
		Logo            string  `json:"logo"`
		Price           float64 `json:"price"`
		CategoryID      uint    `json:"category_id"`
		CycleID         uint    `json:"cycle_id"`
		CurrencyCode    string  `json:"currency"`
		NextPaymentDate string  `json:"next_payment_date"`
	}

	opts := &subscription_handler.HandlerOpts{
		SubscriptionService: service.NewSubscriptionService(memory.NewSubscriptionRepository()),
		CategoryService:     service.NewCategoryService(memory.NewCategoryRepository()),
		CycleService:        service.NewCycleService(memory.NewCycleRepository()),
		CurrencyService:     service.NewCurrencyService(memory.NewCurrencyRepository()),
	}
	ctx := context.Background()

	category1, _ := opts.CategoryService.CreateCategory(ctx, entity.Category{Name: "Test Category 1"})
	category2, _ := opts.CategoryService.CreateCategory(ctx, entity.Category{Name: "Test Category 1"})

	_, _ = opts.CycleService.CreateCycle(ctx, entity.Weekly)
	_, _ = opts.CycleService.CreateCycle(ctx, entity.Monthly)

	_, _ = opts.CurrencyService.CreateCurrency(ctx, entity.RUB)
	_, _ = opts.CurrencyService.CreateCurrency(ctx, entity.USD)

	testCases := []struct {
		name        string
		requestBody req
		expected    resp
		wantErr     error
	}{
		{
			name: "Test Create Subscription 1",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100.0,
				CategoryID:      category1.ID,
				CycleID:         entity.Weekly.ID,
				CurrencyCode:    entity.RUB.Code,
				NextPaymentDate: "2022-01-01",
			},
			expected: resp{
				ID:              1,
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100.0,
				CategoryID:      category1.ID,
				CycleID:         entity.Weekly.ID,
				CurrencyCode:    entity.RUB.Code,
				NextPaymentDate: "2022-01-01",
			},
			wantErr: nil,
		},
		{
			name: "Test Create Subscription 2",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           111.0,
				CategoryID:      category2.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{
				ID:              2,
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           111.0,
				CategoryID:      category2.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			wantErr: nil,
		},
		{
			name: "Test empty name error",
			requestBody: req{
				Name:            "",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           111.0,
				CategoryID:      category2.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  service.ErrInvalidSubscription,
		},
		{
			name: "Test price is zero error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           0,
				CategoryID:      category2.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  service.ErrInvalidSubscription,
		},
		{
			name: "Test price is negative error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           -111.0,
				CategoryID:      category2.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  service.ErrInvalidSubscription,
		},
		{
			name: "Test category not found error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100,
				CategoryID:      10,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  repository.ErrNotFoundCategory,
		},
		{
			name: "Test cycle not found error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100,
				CategoryID:      category1.ID,
				CycleID:         10,
				CurrencyCode:    entity.USD.Code,
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  repository.ErrNotFoundCycle,
		},
		{
			name: "Test currency not found error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100,
				CategoryID:      category1.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    "unknown",
				NextPaymentDate: "2024-05-21",
			},
			expected: resp{},
			wantErr:  repository.ErrNotFoundCurrency,
		},
		{
			name: "Test empty next payment date error",
			requestBody: req{
				Name:            "Test Subscription",
				Note:            "Test Note",
				Logo:            "Test Logo",
				Price:           100,
				CategoryID:      category1.ID,
				CycleID:         entity.Monthly.ID,
				CurrencyCode:    entity.RUB.Code,
				NextPaymentDate: "",
			},
			expected: resp{},
			wantErr:  subscription_handler.ErrInvalidPaymentDate,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			requestBodyBytes, _ := json.Marshal(tc.requestBody)
			r := &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(requestBodyBytes)),
			}

			response := subscription_handler.CreateSubscription(ctx, opts)(r, nil)

			if tc.wantErr != nil {
				assert.ErrorIs(t, tc.wantErr, response.(error))
				return
			}

			tests_assert.EqualAsJSON(t, tc.expected, response)
		})
	}
}
