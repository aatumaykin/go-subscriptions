package currency_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetCurrency(ctx context.Context, cs *service.CurrencyService) api_response.Handle {
	return func(_ *http.Request, ps httprouter.Params) any {
		code := ps.ByName("code")

		currency, err := cs.GetCurrency(ctx, code)
		if err != nil {
			return err
		}

		type resp struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}

		return resp{
			Code:   currency.Code,
			Name:   currency.Name,
			Symbol: currency.Symbol,
		}
	}
}
