package currency_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetCurrencies(ctx context.Context, cs *service.CurrencyService) api_response.Handle {
	return func(_ *http.Request, _ httprouter.Params) any {
		currencies, err := cs.GetAllCurrencies(ctx)
		if err != nil {
			return err
		}

		type resp struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}

		currencyDTOs := make([]resp, len(currencies))
		for i, currency := range currencies {
			currencyDTOs[i] = resp{
				Code:   currency.Code,
				Name:   currency.Name,
				Symbol: currency.Symbol,
			}
		}

		return currencyDTOs
	}
}
