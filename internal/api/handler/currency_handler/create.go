package currency_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/api/handler/error_handler"
	"git.home/alex/go-subscriptions/internal/api/middleware"
	"git.home/alex/go-subscriptions/internal/domain/entity"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func CreateHandle(ctx context.Context, currencyService *service.CurrencyService) httprouter.Handle {
	return middleware.SetJSONContentType(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var requestDTO struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		}
		err := json.NewDecoder(r.Body).Decode(&requestDTO)
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		createdCurrency, err := currencyService.CreateCurrency(ctx, entity.Currency{
			Code:   requestDTO.Code,
			Name:   requestDTO.Name,
			Symbol: requestDTO.Symbol,
		})
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
			Code:   createdCurrency.Code,
			Name:   createdCurrency.Name,
			Symbol: createdCurrency.Symbol,
		}).ToJSON()
		if err != nil {
			error_handler.HandleError(w, err)
			return
		}

		_, _ = w.Write(response)
	})
}
