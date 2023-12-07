package currency_handler

import (
	"context"
	"net/http"

	"git.home/alex/go-subscriptions/internal/api/api_response"
	"git.home/alex/go-subscriptions/internal/domain/service"
	"github.com/julienschmidt/httprouter"
)

func DeleteCurrency(ctx context.Context, cs *service.CurrencyService) api_response.Handle {
	return func(_ *http.Request, ps httprouter.Params) any {
		code := ps.ByName("code")

		err := cs.DeleteCurrency(ctx, code)
		if err != nil {
			return err
		}

		return nil
	}
}
