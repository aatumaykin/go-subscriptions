package currency_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func GetHandle(ctx context.Context, currencyService *service.CurrencyService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
		code := ps.ByName("code")

		currency, err := currencyService.GetCurrency(ctx, code)
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		type responseDTO struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}

		response, err := api_response.Success(responseDTO{
			Code:   currency.Code,
			Name:   currency.Name,
			Symbol: currency.Symbol,
		}).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		_, _ = w.Write(response)
	})
}
