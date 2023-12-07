package currency_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func UpdateCurrency(ctx context.Context, cs *service.CurrencyService) api_response.Handle {
	return func(r *http.Request, ps httprouter.Params) any {
		code := ps.ByName("code")

		var req struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return err
		}

		currency, err := cs.GetCurrency(ctx, code)
		if err != nil {
			return err
		}

		currency.Name = req.Name
		currency.Symbol = req.Symbol

		updatedCurrency, err := cs.UpdateCurrency(ctx, *currency)
		if err != nil {
			return err
		}

		type resp struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}

		return resp{
			Code:   updatedCurrency.Code,
			Name:   updatedCurrency.Name,
			Symbol: updatedCurrency.Symbol,
		}
	}
}
