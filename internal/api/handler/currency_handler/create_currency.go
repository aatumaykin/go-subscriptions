package currency_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CreateCurrency(ctx context.Context, cs *service.CurrencyService) api_response.Handle {
	return func(r *http.Request, _ httprouter.Params) any {
		var req struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		createdCurrency, err := cs.CreateCurrency(ctx, entity.Currency{
			Code:   req.Code,
			Name:   req.Name,
			Symbol: req.Symbol,
		})
		if err != nil {
			return err
		}

		type resp struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}

		return resp{
			Code:   createdCurrency.Code,
			Name:   createdCurrency.Name,
			Symbol: createdCurrency.Symbol,
		}
	}
}
