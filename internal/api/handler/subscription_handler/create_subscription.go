package subscription_handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"github.com/julienschmidt/httprouter"
)

func CreateSubscription(ctx context.Context, ho *HandlerOpts) api_response.Handle {
	return func(r *http.Request, _ httprouter.Params) any {
		var req struct {
			Name            string      `json:"name"`
			Note            string      `json:"note,omitempty"`
			Logo            string      `json:"logo,omitempty"`
			Price           float64     `json:"price"`
			CategoryID      uint        `json:"category_id"`
			CycleID         uint        `json:"cycle_id"`
			CurrencyCode    string      `json:"currency"`
			NextPaymentDate PaymentDate `json:"next_payment_date"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			_, ok := err.(*time.ParseError)
			if ok {
				return ErrInvalidPaymentDate
			}

			return err
		}

		category, err := ho.CategoryService.GetCategory(ctx, req.CategoryID)
		if err != nil {
			return err
		}

		cycle, err := ho.CycleService.GetCycle(ctx, req.CycleID)
		if err != nil {
			return err
		}

		currency, err := ho.CurrencyService.GetCurrency(ctx, req.CurrencyCode)
		if err != nil {
			return err
		}

		createdSubscription, err := ho.SubscriptionService.CreateSubscription(ctx, entity.Subscription{
			Name:            req.Name,
			Note:            req.Note,
			Logo:            req.Logo,
			Price:           req.Price,
			Category:        *category,
			Cycle:           *cycle,
			Currency:        *currency,
			NextPaymentDate: entity.PaymentDate(req.NextPaymentDate),
		})
		if err != nil {
			return err
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

		return resp{
			ID:              createdSubscription.ID,
			Name:            createdSubscription.Name,
			Note:            createdSubscription.Note,
			Logo:            createdSubscription.Logo,
			Price:           createdSubscription.Price,
			CategoryID:      createdSubscription.Category.ID,
			CycleID:         createdSubscription.Cycle.ID,
			CurrencyCode:    createdSubscription.Currency.Code,
			NextPaymentDate: time.Time(createdSubscription.NextPaymentDate).Format(PaymentDateLayout),
		}
	}
}
